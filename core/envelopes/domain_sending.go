package envelopes

import (
	"context"
	"github.com/tietang/dbx"
	"go1234.cn/newResk/core/accounts"
	"go1234.cn/newResk/infra/base"
	"go1234.cn/newResk/services"
	"path"
)

//发红包
func (g *goodsDomain) Send(goods services.RedEnvelopeGoodsDTO) (activity *services.RedEnvelopeActivity, err error) {
	//创建红包
	g.Create(goods)
	//创建活动
	activity = new(services.RedEnvelopeActivity)
	//生成红包链接格式:http://域名/v1/envelope/{id}/link/
	link := base.GetEnvelopeActivityLink()
	domain := base.GetEnvelopeDomain()
	activity.Link = path.Join(domain, link, g.EnvelopeNum)
	//事务
	err = base.Tx(func(runner *dbx.TxRunner) (err error) {
		//将事务与上下文绑定
		ctx := base.WithValueContext(context.Background(), runner)
		//事务中--保存红包
		id, err := g.Save(ctx)
		//保存失败
		if id < 0 || err != nil {

		}

		//红包金额的支付
		//1、红包中间商（配置文件）
		//2、从红包发送人的资金账户扣减红包金额
		//3、将扣减的红包总金额转入红包中间商的红包资金账户

		//交易的主体
		body := services.TradePaticipator{
			AccountNum: goods.AccountNum,
			UserId:     goods.UserId,
			Username:   goods.Username,
		}
		//获取系统红包中间账户信息
		systemAccount := base.GetSystemAccount()
		//交易到的目标
		target := services.TradePaticipator{
			AccountNum: systemAccount.AccountNum,
			UserId:     systemAccount.UserId,
			Username:   systemAccount.Username,
		}
		//创建转账的DTO
		transfer := services.AccountTransferDTO{
			TradeNum:    g.EnvelopeNum, //使用红包的id作为交易的编号
			TradeBody:   body,
			TradeTarget: target,
			Amount:      g.Amount,
			ChangeType:  services.EnvelopeOutGoing,
			ChangeFlag:  services.FlagTransferOut,
			Desc:        "红包金额支付",
		}

		//事务：保存红包商品和红包金额的支付必须保证全部成功或者全部失败，即原子性
		accountDomain := accounts.NewAccountDomain()

		status, err := accountDomain.TransferWithContext(ctx, transfer)
		if status == services.TransferedStatusSuccess {
			return nil
		}

		//在用户账户中成功扣除之后，放入系统红包中间商的账户中，即入账
		transfer = services.AccountTransferDTO{
			TradeNum:    g.EnvelopeNum, //使用红包的id作为交易的编号
			TradeBody:   target,
			TradeTarget: body,
			Amount:      g.Amount,
			ChangeType:  services.EnvelopeInComing,
			ChangeFlag:  services.FlagTransferIn,
			Desc:        "转入系统中间商账户",
		}
		//返回错误，事务会回滚
		status, err = accountDomain.TransferWithContext(ctx, transfer)
		if status == services.TransferedStatusSuccess {
			return nil
		}
		return err
	})
	if err != nil {
		return nil, err
	}

	//扣减成功
	//返回红包链接
	activity.RedEnvelopeGoodsDTO = *g.RedEnvelopeGoods.ToDTO()
	return activity, err
}
