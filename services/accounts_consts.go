package services

//转账状态
type TransferedStatus int8

const (
	//转账失败
	TransferedStatusFailure TransferedStatus = -1
	//余额不足
	TransferedStatusSufficientFunds TransferedStatus = 0
	//转账成功
	TransferedStatusSuccess TransferedStatus = 1
)

//转账的类型 ：0=创建账户 >=1进账 <=-1支出
type ChangeType int8

const (
	//创建账户
	AccountCreated ChangeType = 0
	//储值
	AccountStoreValue ChangeType = 1
	//红包资金的支出
	EnvelopeOutGoing ChangeType = -2
	//红包的收入
	EnvelopeInComing ChangeType = 2
	//红包的过期退款
	EnvelopeExpiredRefund ChangeType = 3
)

//资金交易的变化标识 创建账户0 支出-1 收入1
type ChangeFlag int8

const (
	FlagAccountCreated ChangeFlag = 0
	FlagTransferOut    ChangeFlag = 1
	FlagTransferIn     ChangeFlag = -1
)
