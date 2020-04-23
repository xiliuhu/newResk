package envelopes

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	acservices "github.com/ztaoing/account/services"
	"github.com/ztaoing/infra/base"
	"github.com/ztaoing/newResk/services"
)

const pageSize = 100

//红包过期
type ExpiredEnvelopeDomain struct {
	ExpiredGoods []RedEnvelopeGoods
	offset       int
}

//按照分页查出指定数量的过期红包
func (e *ExpiredEnvelopeDomain) Next() (ok bool) {
	base.Tx(func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeGoodsDao{runner: runner}
		e.ExpiredGoods = dao.FindExpired(e.offset, pageSize)
		//查询出结果
		if len(e.ExpiredGoods) > 0 {
			e.offset += len(e.ExpiredGoods)
			ok = true
		}
		return nil
	})
	return ok
}

//在定时任务中调用
func (e *ExpiredEnvelopeDomain) Expired() (err error) {
	for e.Next() {
		for _, g := range e.ExpiredGoods {
			log.Debugf("过期红包对款开始：%v", g)
			//对单个红包触发退款流程
			err = e.ExpiredOne(g)
			if err != nil {
				log.Error(err)
			}
			log.Debugf("过期退款结束:%v", g)
		}
	}
	return nil
}

//对单个红包触发退款流程
func (e *ExpiredEnvelopeDomain) ExpiredOne(goods RedEnvelopeGoods) (err error) {
	//创建退款订单(在原订单的基础上进行修改)
	refund := goods
	refund.OrderType = services.OrderTypeRefund
	refund.RemainAmount = goods.RemainAmount.Mul(decimal.NewFromFloat(-1))
	refund.RemainQuantity = -goods.RemainQuantity
	//订单状态定义为过期
	refund.Status = services.OrderExpire
	//支付状态
	refund.PayStatus = services.Refunding
	//原订单号
	refund.OriginEnvelopeNum = goods.EnvelopeNum
	domain := goodsDomain{RedEnvelopeGoods: refund}
	//生成新的订单号
	domain.createEnvelopeNum()

	//在事务中进行修改
	err = base.Tx(func(runner *dbx.TxRunner) error {
		txCTX := base.WithValueContext(context.Background(), runner)
		//创建退款订单
		id, err := domain.Save(txCTX)
		if err != nil || id == 0 {
			return errors.New("创建退款订单失败")
		}
		//创建退款订单成功后，修改原订单的状态
		dao := RedEnvelopeGoodsDao{runner: runner}
		rows, err := dao.UpdateOrderStatus(goods.EnvelopeNum, services.OrderExpire)
		if err != nil || rows == 0 {
			return errors.New("更新原订单状态失败")
		}

		return nil
	})
	if err != nil {
		return
	}
	//调用资金账户接口转账，将剩余金额退回给用户
	systemAccount := base.GetSystemAccount()
	//红包所有者的账户信息
	account := acservices.GetAccountService().GetEnvelopeAccountByUserId(goods.UserId)
	if account == nil {
		return errors.New("没有找到此红包所有者的资金账户:" + goods.UserId)
	}
	//交易主体
	body := acservices.TradePaticipator{
		AccountNum: systemAccount.AccountNum,
		UserId:     systemAccount.UserId,
		Username:   systemAccount.Username,
	}
	//交易对象
	target := acservices.TradePaticipator{
		AccountNum: account.AccountNum,
		UserId:     account.UserId,
		Username:   account.Username,
	}
	//创建转账的DTO
	transfer := acservices.AccountTransferDTO{
		TradeNum:    refund.EnvelopeNum,
		TradeBody:   body,
		TradeTarget: target,
		Amount:      goods.RemainAmount,
		ChangeType:  acservices.EnvelopeExpiredRefund,
		ChangeFlag:  acservices.FlagTransferIn,
		Desc:        "红包过期退款:" + goods.EnvelopeNum,
	}
	//执行转账
	status, err := acservices.GetAccountService().Transfer(transfer)
	if status != acservices.TransferedStatusSuccess {
		return err
	}

	err = base.Tx(func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeGoodsDao{runner: runner}
		//修改原订单状态
		rows, err := dao.UpdateOrderStatus(goods.EnvelopeNum, services.OrderExpireRefundSucessful)
		if err != nil || rows == 0 {
			return errors.New("修改原订单状态失败")
		}
		//修改退款订单状态
		rows, err = dao.UpdateOrderStatus(refund.EnvelopeNum, services.OrderExpireRefundSucessful)
		if err != nil || rows == 0 {
			return errors.New("修改原订单状态失败")
		}
		return nil
	})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
