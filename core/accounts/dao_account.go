package accounts

import (
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type AccountDao struct {
	runner *dbx.TxRunner //事务
}

//查询数据库持久化对象 单实例，获取一行数据
func (dao *AccountDao) GetOne(accountNo string) *Account {
	a := &Account{
		AccountNo: accountNo,
	}
	ok, err := dao.runner.GetOne(a)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	//数据不存在
	if !ok {
		return nil
	}
	return a
}

//资金账户
//一个用户可以拥有多个类型的资金账户
//通过用户id和账户类型来查询账户信息
func (dao *AccountDao) GetByUserId(userId string, accountType int) *Account {
	sql := "select * from account " +
		"where user_id=? and account_type=? "
	a := &Account{}
	ok, err := dao.runner.Get(a, sql, userId, accountType)
	if err != nil {
		logrus.Error(err)
		return nil
	}
	//数据不存在
	if !ok {
		return nil
	}
	return a
}

//账户数据的插入
func (dao *AccountDao) Insert(a *Account) (id int64, err error) {
	res, err := dao.runner.Insert(a)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

//账户数据的更新
func (dao *AccountDao) Update(accountNo string, amount decimal.Decimal) (rows int64, err error) {
	//此处使用了乐观锁
	sql := "update account" +
		"set balance=balance+CAST(? AS DECIMAL(30,6))" +
		"where account_no=? " +
		"and balance>=-1*CAST(? AS DECIMAL(30,6))"
	res, err := dao.runner.Exec(sql, amount.String(), accountNo)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}

//账户状态更新
func (dao *AccountDao) UpdateAccount(accountNo int64, status int) (rows int64, err error) {
	sql := "update account" +
		"set status=? " +
		"where account_no=?"
	res, err := dao.runner.Exec(sql, status, accountNo)
	if err != nil {
		return 0, err
	}
	return res.RowsAffected()
}
