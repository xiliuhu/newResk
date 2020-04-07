package accounts

import (
	"errors"
	"github.com/segmentio/ksuid"
	"github.com/tietang/dbx"
	"go1234.cn/newResk/infra/base"
	"go1234.cn/newResk/services"
)

//领域模型是有状态的，每次使用时都要实例化
//只能在accounts包中使用
type accountDomain struct {
	account    Account
	accountLog AccountLog
}

//创建流水记录
func (domain *accountDomain) creatAccountLog() {
	domain.accountLog = AccountLog{}
	domain.createAccountLogNo()
	domain.accountLog.TradeNo = domain.accountLog.AccountNo

	//流水中的交易主体信息
	domain.accountLog.AccountNo = domain.account.AccountNo
	domain.accountLog.UserId = domain.account.UserId
	domain.accountLog.Username = domain.account.Username.String

	//交易对象信息
	domain.accountLog.TargetAccountNo = domain.account.AccountNo
	domain.accountLog.TargetUserId = domain.account.UserId
	domain.accountLog.TargetUsername = domain.account.Username.String

	//交易金额
	domain.accountLog.Amount = domain.account.Balance  //交易的金额
	domain.accountLog.Balance = domain.account.Balance //交易之后的余额

	//交易变化属性
	domain.accountLog.Decs = "创建账户"
	domain.accountLog.ChangeType = services.AccountCreated
	domain.accountLog.ChangeFlag = services.FlagAccountCreated
}

//创建账户
func (domain *accountDomain) CreateAccount(dto services.AccountDTO) (*services.AccountDTO, error) {
	//创建账户持久化对象
	domain.account = Account{}
	domain.account.FromDTO(&dto) //转换
	domain.createAccountNo()
	domain.account.Username.Valid = true //为true时才向数据库写入

	//创建账户流水的持久化对象
	domain.createAccountLogNo()
	accountDao := AccountDao{}
	accountLogDao := AccountLogDao{}
	var rdto *services.AccountDTO
	//快捷的事务函数，返回为非nil则会回滚
	err := base.Tx(func(runner *dbx.TxRunner) error {
		accountDao.runner = runner
		accountLogDao.runner = runner
		//插入账户数据，然后插入流水数据
		id, err := accountDao.Insert(&domain.account)
		if err != nil {
			return nil
		}
		if id <= 0 {
			return errors.New("创建账户失败")
		}
		//插入流水数据
		id, err = accountLogDao.Insert(&domain.accountLog)
		if err != nil {
			return nil
		}
		if id <= 0 {
			return errors.New("创建流水失败")
		}
		//通过账户编号查出账户信息
		domain.account = *accountDao.GetOne(domain.account.AccountNo)
		return nil
	})
	//转换成DTO对象
	rdto = domain.account.ToDTO()
	return rdto, err
}

//创建流水logNo
func (domain *accountDomain) createAccountLogNo() {
	//暂时使用ksuid生成ID
	//后期需要使用优化成分布式ID
	//全局唯一的ID
	domain.accountLog.LogNo = ksuid.New().Next().String()
}

//生成账户accountNo
func (domain *accountDomain) createAccountNo() {
	domain.account.AccountNo = ksuid.New().Next().String()
}
