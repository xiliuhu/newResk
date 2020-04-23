package envelopes

import (
	"database/sql"
	"github.com/shopspring/decimal"
	"github.com/ztaoing/newResk/services"
	"time"
)

type RedEnvelopeItem struct {
	Id           int64              `db:"id,omitempty"`  //自增ID
	ItemNum      string             `db:"item_no,uni"`   //红包订单详情编号
	EnvelopeNum  string             `db:"envelope_no"`   //红包编号
	RecvUsername sql.NullString     `db:"recv_user_id"`  //接收红包的用户名
	RecvUserId   string             `db:"recv_user_id"`  //接收红包的用户ID
	Amount       decimal.Decimal    `db:"amount"`        //红包金额
	Quantity     int                `db:"quantity"`      //收红包的数量
	RemainAmount decimal.Decimal    `db:"remain_amount"` //红包剩余金额
	AccountNum   string             `db:"account_no"`    //接收红包账户编号
	PayStatus    services.PayStatus `db:"pay_status"`    //支付状态：未支付，支付中，已支付，支付失败
	CreatedAt    time.Time          `db:"created_at"`    //创建时间
	UpdatedAt    time.Time          `db:"updated_at"`    //更新时间
	Desc         string             `db:"desc"`          //描述
}

func (po *RedEnvelopeItem) ToDTO() *services.RedEnvelopeItemDTO {
	dto := &services.RedEnvelopeItemDTO{

		ItemNum:      po.ItemNum,
		EnvelopeNum:  po.EnvelopeNum,
		RecvUsername: po.RecvUsername.String,
		RecvUserId:   po.RecvUserId,
		Amount:       po.Amount,
		Quantity:     po.Quantity,
		RemainAmount: po.RemainAmount,
		AccountNum:   po.AccountNum,
		PayStatus:    po.PayStatus,
		CreatedAt:    po.CreatedAt,
		UpdatedAt:    po.UpdatedAt,
		Desc:         po.Desc,
	}
	return dto
}

func (po *RedEnvelopeItem) FromDTO(dto *services.RedEnvelopeItemDTO) {

	po.ItemNum = dto.ItemNum
	po.EnvelopeNum = dto.EnvelopeNum
	po.RecvUsername = sql.NullString{Valid: true, String: dto.RecvUsername}
	po.RecvUserId = dto.RecvUserId
	po.Amount = dto.Amount
	po.Quantity = dto.Quantity
	po.RemainAmount = dto.RemainAmount
	po.AccountNum = dto.AccountNum
	po.PayStatus = dto.PayStatus
	po.CreatedAt = dto.CreatedAt
	po.UpdatedAt = dto.UpdatedAt
	po.Desc = dto.Desc
}
