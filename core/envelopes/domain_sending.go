package envelopes

import (
	"go1234.cn/newResk/infra/base"
	"go1234.cn/newResk/services"
	"path"
)

//发红包
func (g *goodsDomain) Send(goods services.RedEnvelopeGoodsDTO) (activity *services.RedEnvelopeActivity) {
	//创建红包
	g.Create(goods)
	//创建活动
	activity = new(services.RedEnvelopeActivity)
	//保存红包
	link := base.GetEnvelopeActivityLink()
	domain := base.GetEnvelopeDomain()
	activity.Link = path.Join(domain, link, g.EnvelopeNum)
	//红包金额的支付
	//1、红包中间商（配置文件）
	//2、从红包发送人的资金账户扣减红包金额
	//3、将扣减的红包总金额转入红包中间商的红包资金账户

	//扣减成功
	return
}
