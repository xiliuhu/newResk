package envelopes

import (
	"context"
	"database/sql"
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/ztaoing/account/core/accounts"
	acservices "github.com/ztaoing/account/services"
	"github.com/ztaoing/infra/algo"
	"github.com/ztaoing/infra/base"
	"github.com/ztaoing/newResk/services"
)

//金额转换
var multiple = decimal.NewFromFloat(100.0)

//抢红包的业务领域层
func (g *goodsDomain) Receive(ctx context.Context, dto services.RedEnvelopeReceiveDTO) (item *services.RedEnvelopeItemDTO, err error) {
	//创建收红包的订单
	g.preCreateItem(dto)
	//查询当前红包的剩余数量和剩余金额信息
	goods := g.Get(dto.EnvelopeNum)
	//校验当前红包的剩余数量和剩余金额：如果没有剩余，直接返回无可用红包金额
	if goods.RemainQuantity <= 0 || goods.RemainAmount.Cmp(decimal.NewFromFloat(0)) <= 0 {
		return nil, errors.New("剩余金额不足")
	}
	//使用红包算法计算红包金额
	nextAmount := g.nextAmount(goods)
	//使用乐观锁更新语句，尝试更新剩余数量和剩余金额：
	err = base.Tx(func(runner *dbx.TxRunner) error {

		dao := RedEnvelopeGoodsDao{runner: runner}
		rows, err := dao.UpdateBalance(g.EnvelopeNum, nextAmount)

		//-更新失败，抢红包失败，没有可用的红包数量和金额，返回0
		if rows <= 0 || err != nil {
			return errors.New("更新失败，红包不足")
		}
		//-更新成功，抢到红包，返回1
		//保存订单明细数据
		g.item.Quantity = 1
		g.item.PayStatus = services.Paying
		g.item.AccountNum = dto.AccountNum
		g.item.RemainAmount = goods.RemainAmount.Sub(nextAmount)
		g.item.Amount = nextAmount
		txCTX := base.WithValueContext(ctx, runner)
		_, err = g.item.Save(txCTX)
		if err != nil {
			log.Error(err)
			return err
		}
		//将抢到的红包从系统红包中间账户转入，抢到红包的用户
		status, err := g.transfer(txCTX, dto)
		if status == acservices.TransferedStatusSuccess {
			return nil
		}
		//转账失败
		return err

	})
	return nil, err

}

//创建收红包的订单
func (g *goodsDomain) preCreateItem(dto services.RedEnvelopeReceiveDTO) {
	g.item.AccountNum = dto.AccountNum
	g.item.EnvelopeNum = dto.EnvelopeNum
	g.item.RecvUsername = sql.NullString{String: dto.RecvUsername, Valid: true}
	g.item.RecvUserId = dto.RecvUserId
	g.item.createItemNum()

}

//使用二倍平均算法计算红包金额
func (g *goodsDomain) nextAmount(goods *RedEnvelopeGoods) (amount decimal.Decimal) {
	//剩余红包数量为1，返回剩余金额
	if g.RemainQuantity == 1 {
		return goods.RemainAmount
	}
	//如果是等额红包
	if goods.EnvelopeType == services.GeneralEnvelope {
		return goods.AmountOne
	}
	//如果是碰运气红包
	if goods.EnvelopeType == services.LuckyEnvelope {
		//分 = 元*100，并取出int值
		cents := goods.RemainAmount.Mul(multiple).IntPart()
		//使用二倍平均算法
		next := algo.DoubleAverage(int64(g.RemainQuantity), cents)
		//由分转化为元
		amount = decimal.NewFromFloat(float64(next)).Div(multiple)

	}
	return amount
}

//由系统中间账户 转入到 用户账户
func (g *goodsDomain) transfer(ctx context.Context, dto services.RedEnvelopeReceiveDTO) (status acservices.TransferedStatus, err error) {
	systemAccount := base.GetSystemAccount()
	body := acservices.TradePaticipator{
		AccountNum: systemAccount.AccountNum,
		UserId:     systemAccount.UserId,
		Username:   systemAccount.AccountName,
	}
	target := acservices.TradePaticipator{
		AccountNum: dto.AccountNum,
		UserId:     dto.RecvUserId,
		Username:   dto.RecvUsername,
	}
	transfer := acservices.AccountTransferDTO{
		TradeNum:    dto.EnvelopeNum,
		TradeBody:   body,
		TradeTarget: target,
		Amount:      g.item.Amount,
		ChangeType:  acservices.EnvelopeInComing,
		ChangeFlag:  acservices.FlagTransferIn,
		Desc:        "红包转入",
	}
	domain := accounts.NewAccountDomain()
	return domain.TransferWithContext(ctx, transfer)
}
