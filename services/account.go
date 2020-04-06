package services

import "time"

type AccountService interface {
	CreateAccount(dto AccountCreateDTO) (*CreateDTO, error) //创建账户接口
	Transfer()                                              //转账接口
	StoreValue()                                            //储值接口
	GetEnvelopeAccountByUserId()                            //查询账户接口
}

//CreateAccount 入参
type AccountCreateDTO struct {
	UserId       string //用户id
	UserName     string //用户名称
	AccountName  string //账户名称
	AccountType  string //账户的类型
	CurrencyCode string //
	Amount       string //金额 使用string，防止在传递的过程中丢失精度
}

//CreateAccount 出参
type CreateDTO struct {
	AccountCreateDTO
	AccountNo string    //账户编号
	CreateAt  time.Time //创建时间
}

//Transfer接口的入参
//交易的参与者
type TradePaticipator struct {
	AccountNo string //交易变化
	UserId    string //交易的id
	UserName  string //交易的名称

}
