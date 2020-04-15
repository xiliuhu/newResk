package services

const (
	DefaultBlessing   = "恭喜发财啊！亲"
	DefaultTimeFormat = "2006-01-02.15:04:05"
)

//订单类型
type OrderType int

const (
	OrderTypeSending OrderType = 1
	OrderTypeRefund  OrderType = 2
)

//支出状态:未支付，支出中，已支付，支付失败
//退款：未退款，退款中，已退款，退款失败
type PayStatus int

const (
	UnPay         PayStatus = 1
	Paying        PayStatus = 2
	Payed         PayStatus = 3
	PayFailure    PayStatus = 4
	UnRefund      PayStatus = 61
	Refunding     PayStatus = 62
	Refunded      PayStatus = 63
	RefundFailure PayStatus = 64
)

//红包订单状态：创建、发布、过期、失效、过期退款成功、过期退款失败
type OrderStatus int

const (
	OrderCreate                OrderStatus = 1
	OrderSending               OrderStatus = 2
	OrderExpire                OrderStatus = 3
	OrderDisabled              OrderStatus = 4
	OrderExpireRefundSucessful OrderStatus = 5
	OrderExpireRefundFailure   OrderStatus = 6
)

//红包的类型：普通红包、碰运气红包
type EnvelopeType int

const (
	GeneralEnvelope EnvelopeType = 1
	LuckyEnvelope   EnvelopeType = 2
)

var EnvelopeTypes = map[EnvelopeType]string{
	GeneralEnvelope: "普通红包",
	LuckyEnvelope:   "碰运气红包",
}
