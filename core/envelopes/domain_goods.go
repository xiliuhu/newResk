package envelopes

import (
	"github.com/segmentio/ksuid"
	"github.com/shopspring/decimal"
	"github.com/tietang/dbx"
	"go1234.cn/newResk/infra/base"
	"go1234.cn/newResk/services"
	"time"
)

type goodsDomain struct {
	RedEnvelopeGoods
}

//1、生成一个红包编号
func (g *goodsDomain) createEnvelopeNum() {
	g.EnvelopeNum = ksuid.New().Next().String()

}

//创建一个红包商品对象
func (g *goodsDomain) Create(goods services.RedEnvelopeGoodsDTO) {
	g.RedEnvelopeGoods.FromDTO(&goods)

	g.RemainQuantity = goods.Quantity
	g.Username.Valid = true
	g.Blessing.Valid = true
	//普通红包
	if g.EnvelopeType == services.GeneralEnvelope {
		g.Amount = goods.AmountOne.Mul(decimal.NewFromFloat(float64(goods.Quantity)))
	}
	//碰运气红包
	if g.EnvelopeType == services.LuckyEnvelope {
		g.AmountOne = decimal.NewFromFloat(0)
	}
	g.RemainAmount = goods.Amount
	//过期时间
	g.ExpiredAt = time.Now().Add(24 * time.Hour)
	g.Status = services.OrderCreate
	//生成红包编号
	g.createEnvelopeNum()
}

//保存到红包商品表
func (g *goodsDomain) Save(goods services.RedEnvelopeGoodsDTO) (id int64, err error) {
	err = base.TX(func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeGoodsDao{runner: runner}
		id, err = dao.Insert(&g.RedEnvelopeGoods)
		return err
	})
	return id, err
}

//创建并保存红包商品
func (g *goodsDomain) CreateAndSave(goods services.RedEnvelopeGoodsDTO) (id int64, err error) {
	//创建红包商品
	g.Create(goods)
	//保存红包商品
	g.Save(goods)

}
