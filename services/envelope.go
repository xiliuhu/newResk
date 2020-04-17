package services

import (
	"github.com/shopspring/decimal"
	"go1234.cn/newResk/infra/base"
	"time"
)

var IRedEnvelopeService RedEnvelopeService

func GetRedEnvelopeService() RedEnvelopeService {
	base.Check(IRedEnvelopeService)
	return IRedEnvelopeService
}

type RedEnvelopeService interface {
	//发红包
	Send(dto RedEnvelopeSendDTO) (activity *RedEnvelopeActivity, err error)
	//收红包
	Recive(dto RedEnvelopeReceiveDTO) (item *RedEnvelopeItemDTO, err error)
	//退款
	Refund(envelopeNum string) (order *RedEnvelopeGoodsDTO)
	//查询红包订单
	Get(envelopeNum string) (order *RedEnvelopeGoodsDTO)
	//查询用户已经发送的红包列表
	ListSend(userId string, page, size int) (order []*RedEnvelopeGoodsDTO)
	ListReceived(userId string, page, size int) (items []*RedEnvelopeItemDTO)
	//查询用户已经抢到的红包列表
	ListReceivable(page, size int) (orders []*RedEnvelopeGoodsDTO)
	ListItems(envelopeNum string) (items []*RedEnvelopeItemDTO)
}

//发送红包
type RedEnvelopeSendDTO struct {
	EnvelopeType EnvelopeType    `json:"envelope_type" validate:"required"`    //红包类型：普通红包，碰运气红包
	Username     string          `json:"user_name",validate:"required"'`       //用户名称
	UserId       string          `json:"user_id" validate:"required"`          //发送红包的用户ID
	Blessing     string          `json:"blessing"`                             //祝福语
	Amount       decimal.Decimal `json:"amount" validate:"required,numeric"`   //红包金额：等额红包、碰运气红包
	Quantity     int             `json:"quantity" validate:"required,numeric"` //红包总数
}

type RedEnvelopeActivity struct {
	RedEnvelopeGoodsDTO
	Link string `json:"link"` //红包链接
}

func (r *RedEnvelopeActivity) CopyTo(target *RedEnvelopeActivity) {
	target.Link = r.Link
	target.EnvelopeNum = r.EnvelopeNum
	target.EnvelopeType = r.EnvelopeType
	target.Username = r.Username
	target.UserId = r.UserId
	target.Blessing = r.Blessing
	target.Amount = r.Amount
	target.AmountOne = r.AmountOne
	target.Quantity = r.Quantity
	target.RemainAmount = r.RemainAmount
	target.RemainQuantity = r.RemainQuantity
	target.ExpiredAt = r.ExpiredAt
	target.Status = r.Status
	target.OrderType = r.OrderType
	target.PayStatus = r.PayStatus
	target.CreatedAt = r.CreatedAt
	target.UpdatedAt = r.UpdatedAt
}

//将实例转化为GoodsDTO
func (r *RedEnvelopeSendDTO) ToGoods() *RedEnvelopeGoodsDTO {
	goods := &RedEnvelopeGoodsDTO{
		EnvelopeType: r.EnvelopeType,
		Username:     r.Username,
		UserId:       r.UserId,
		Blessing:     r.Blessing,
		Amount:       r.Amount,
		Quantity:     r.Quantity,
	}
	return goods
}

//接收红包
type RedEnvelopeReceiveDTO struct {
	EnvelopeNum  string `json:"envelope_no",validate:"required"`  //红包编号
	RecvUsername string `json:"recv_username",valiate:"required"` //红包接收者的名称
	RecvUserId   string `json:"recv_user_id",validate:"required"` //红包接收者的用户ID
	AccountNum   string `json:"account_no"`                       //账户编号
}

//红包
type RedEnvelopeGoodsDTO struct {
	EnvelopeNum       string          `json:"envelope_no"`                            //红包编号，红包唯一标识
	EnvelopeType      EnvelopeType    `json:"envelope_type",validate:"required"`      //红包类型：普通红包、碰运气红包
	Username          string          `json:"username",validate:"required"`           //用户名称
	UserId            string          `json:"user_id",validate:"required"`            //用户ID
	Blessing          string          `json:"blessing"`                               //祝福语
	Amount            decimal.Decimal `json:"amount",validate:"required,numeric"`     //红包总金额
	AmountOne         decimal.Decimal `json:"amount_one",validate:"required,numeric"` //单个红包的金额，碰运气红包无
	Quantity          int             `json:"quantity",validate:"required,numeric"`   //红包总数量
	RemainAmount      decimal.Decimal `json:"remain_amount"`                          //剩余红包金额
	RemainQuantity    int             `json:"remain_quantity"`                        //剩余红包数量
	ExpiredAt         time.Time       `json:"expired_at"`                             //过期时间
	Status            OrderStatus     `json:"status"`                                 //红包状态
	OrderType         OrderType       `json:"order_type"`                             //订单类型：发布订单，退款单
	PayStatus         PayStatus       `json:"pay_status"`                             //支付状态：未支付、支付中、已支付、支付失败
	CreatedAt         time.Time       `json:"created_at"`                             //创建时间
	UpdatedAt         time.Time       `json:"update_at"`                              //更新时间
	AccountNum        string          `json:"account_no"`                             //账户编号
	OriginEnvelopeNum string          `json:"origin_envelope_no"`
}

//红包详情
type RedEnvelopeItemDTO struct {
	ItemNum      string          `json:"item_no"`       //红包订单详情编号
	EnvelopeNum  string          `json:"envelope_no"`   //红包订单编号
	RecvUsername string          `json:"recv_username"` //红包接受者用户名称
	RecvUserId   string          `json:"recv_user_id"`  //红包接受者用户ID
	Amount       decimal.Decimal `json:"amount"`        //金额
	Quantity     int             `json:"quantity"`      //收到的数量：
	RemainAmount decimal.Decimal `json:"remain_amount"` //收到红包剩余金额
	AccountNum   string          `json:"account_no"`    //红包接受者账户
	PayStatus    PayStatus       `json:"pay_status"`    //支付状态：未支付、支付中、已支付、支付失败
	CreatedAt    time.Time       `json:"created_at"`    //创建时间
	UpdatedAt    time.Time       `json:"updated_at"`    //更新时间
	IsLucy       bool            `json:"is_lucy"`       //是否为幸运红包
	Desc         string          `json:"desc"`          //描述
}

func (r RedEnvelopeItemDTO) CopyTo(item *RedEnvelopeItemDTO) {
	item.ItemNum = r.ItemNum
	item.EnvelopeNum = r.EnvelopeNum
	item.RecvUsername = r.RecvUsername
	item.RecvUserId = r.RecvUserId
	item.Amount = r.Amount
	item.Quantity = r.Quantity
	item.RemainAmount = r.RemainAmount
	item.AccountNum = r.AccountNum
	item.PayStatus = r.PayStatus
	item.CreatedAt = r.CreatedAt
	item.UpdatedAt = r.UpdatedAt
	item.Desc = r.Desc
}
