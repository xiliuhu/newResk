package envelopes

import (
	"context"
	"errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"github.com/ztaoing/infra/base"
	"github.com/ztaoing/newResk/services"
	"sync"
)

//发送红包的应用服务层
var once sync.Once

func init() {
	once.Do(func() {
		services.IRedEnvelopeService = new(redEnvelopeService)
	})
}

type redEnvelopeService struct {
}

//发送红包
func (*redEnvelopeService) Send(dto services.RedEnvelopeSendDTO) (activity *services.RedEnvelopeActivity, err error) {
	//验证
	if err = base.ValidateStructs(&dto); err != nil {
		return activity, err
	}

	//获取红包发送者的资金账户信息
	account := services.GetAccountService().GetEnvelopeAccountByUserId(dto.UserId)
	if account != nil {
		return nil, errors.New("此用户不存在:Id-" + dto.UserId)
	}

	goods := dto.ToGoods()
	goods.AccountNum = account.AccountNum

	//祝福语
	if goods.Blessing == "" {
		goods.Blessing = services.DefaultBlessing
	}
	//等额类型的红包
	if goods.EnvelopeType == services.GeneralEnvelope {
		goods.AmountOne = goods.Amount
		goods.Amount = decimal.Decimal{}
	}
	//goodsDomain是有状态的所以使用的时候每次都需要new
	domain := new(goodsDomain)
	//执行发送红包的逻辑
	activity, err = domain.Send(*goods)
	if err != nil {
		//将错误打印到控制台
		log.Error(err)
	}
	return activity, err
}

//收红包
func (r *redEnvelopeService) Receive(dto services.RedEnvelopeReceiveDTO) (item *services.RedEnvelopeItemDTO, err error) {
	//参数验证
	if err = base.ValidateStructs(&dto); err != nil {
		return nil, err
	}
	//获取当前红包账户信息
	account := services.GetAccountService().GetEnvelopeAccountByUserId(dto.RecvUserId)
	if account != nil {
		return nil, errors.New("红包资金账户不存在：" + dto.RecvUserId)
	}
	dto.AccountNum = account.AccountNum
	//收红包
	domain := goodsDomain{}
	itemDomain := itemDomain{}
	//获取红包订单详情
	item = itemDomain.GetByUser(dto.RecvUserId, dto.EnvelopeNum)
	if item != nil {
		return item, nil
	}
	item, err = domain.Receive(context.Background(), dto)
	return item, err
}
