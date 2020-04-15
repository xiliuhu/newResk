package envelopes

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"go1234.cn/newResk/services"
	"time"
)

type RedEnvelopeGoods struct {
	Id                int64                 `db:"id,omitempty"`         //自增ID
	EnvelopeNum       string                `db:"envelope_no,uni"`      //红包编号，红包唯一标识
	EnvelopeType      services.EnvelopeType `db:"envelope_type"`        //红包类型：普通红包、碰运气红包
	Username          sql.NullString        `db:"username"`             //用户名称
	UserId            string                `db:"user_id"`              //用户ID
	Blessing          sql.NullString        `db:"blessing""`            //祝福语
	Amount            decimal.Decimal       `db:"amount"`               //红包总金额
	AmountOne         decimal.Decimal       `db:"amount_one"`           //单个红包的金额
	Quantity          int                   `db:"quantity"`             //红包总数量
	RemainAmount      decimal.Decimal       `db:"remain_amount"`        //剩余红包金额
	RemainQuantity    int                   `db:"remain_quantity"`      //剩余红包数量
	ExpiredAt         time.Time             `db:"expired_at"`           //过期时间
	Status            services.OrderStatus  `db:"status"`               //红包状态：0初始化 1启用 2失效
	OrderType         services.OrderType    `db:"order_type"`           //支付类型：发布、退款
	PayStatus         services.PayStatus    `db:"pay_status"`           //支付状态：未支付、支付中、已支付、支付失败
	CreatedAt         time.Time             `db:"create_at,omitempty"`  //创建时间
	UpdatedAt         time.Time             `db:"updated_at,omitempty"` //更新时间
	OriginEnvelopeNum string                `db:"origin_envelope"`      //原关联订单号

}

func (po *RedEnvelopeGoods) ToDTO() *services.RedEnvelopeGoodsDTO {
	dto := &services.RedEnvelopeGoodsDTO{

		EnvelopeNum:       po.EnvelopeNum,
		EnvelopeType:      po.EnvelopeType,
		Username:          po.Username.String,
		UserId:            po.UserId,
		Blessing:          po.Blessing.String,
		Amount:            po.Amount,
		AmountOne:         po.AmountOne,
		Quantity:          po.Quantity,
		RemainAmount:      po.RemainAmount,
		RemainQuantity:    po.RemainQuantity,
		ExpiredAt:         po.ExpiredAt,
		Status:            po.Status,
		OrderType:         po.OrderType,
		PayStatus:         po.PayStatus,
		CreatedAt:         po.CreatedAt,
		UpdatedAt:         po.UpdatedAt,
		OriginEnvelopeNum: po.OriginEnvelopeNum,
	}
	return dto
}

func (po *RedEnvelopeGoods) FromDTO(dto *services.RedEnvelopeGoodsDTO) {
	po.EnvelopeNum = dto.EnvelopeNum
	po.EnvelopeType = dto.EnvelopeType
	po.Username = sql.NullString{Valid: true, String: dto.Username}
	po.UserId = dto.UserId
	po.Blessing = sql.NullString{Valid: true, String: dto.Blessing}
	po.Amount = dto.Amount
	po.AmountOne = dto.AmountOne
	po.Quantity = dto.Quantity
	po.RemainAmount = dto.RemainAmount
	po.RemainQuantity = dto.RemainQuantity
	po.ExpiredAt = dto.ExpiredAt
	po.Status = dto.Status
	po.OrderType = dto.OrderType
	po.PayStatus = dto.PayStatus
	po.CreatedAt = dto.CreatedAt
	po.UpdatedAt = dto.UpdatedAt
	po.OriginEnvelopeNum = dto.OriginEnvelopeNum
}
