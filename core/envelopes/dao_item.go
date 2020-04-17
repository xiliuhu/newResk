package envelopes

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
)

type RedEnvelopeItemDao struct {
	runner *dbx.TxRunner
}

//查询
func (dao *RedEnvelopeItemDao) GetOne(itemNo string) *RedEnvelopeItem {
	form := &RedEnvelopeItem{ItemNo: itemNo}
	ok, err := dao.runner.GetOne(form)
	if err != nil {
		return nil
	}
	if !ok {
		return nil
	}
	return form
}

//红包订单详情数据的写入 Insert
func (dao *RedEnvelopeItemDao) Insert(form *RedEnvelopeItem) (int64, error) {
	rs, err := dao.runner.Insert(form)
	if err != nil {
		return 0, err
	}
	return rs.LastInsertId()
}

//查询红包列表-红包编号
func (dao *RedEnvelopeItemDao) FindItems(envelopeNo string) []*RedEnvelopeItem {
	items := make([]*RedEnvelopeItem, 0)
	sql := "select * from red_envelope_item where envelope_no=?"
	err := dao.runner.Find(&items, sql, envelopeNo)
	if err != nil {
		logrus.Error(err)
		return items
	}
	return items
}

//查询红包记录
func (dao *RedEnvelopeItemDao) GetByUser(envelopeNo, userId string) *RedEnvelopeItem {
	item := RedEnvelopeItem{}
	sql := "select * from red_envelope_item where envelope_no=? and recv_user_id=?"
	ok, err := dao.runner.Get(&item, sql, envelopeNo, userId)
	if !ok {
		return nil
	}
	if err != nil {
		logrus.Error(err)
		return nil
	}
	return &item
}

func (dao *RedEnvelopeItemDao) ListReceivedItems(userId string, page, size int) []*RedEnvelopeItem {
	items := make([]*RedEnvelopeItem, 0)
	sql := "select * from red_envelope_item where recv_user_id=? order by created_at desc limit ?,?"
	err := dao.runner.Find(&items, sql, userId, page, size)
	if err != nil {
		logrus.Error(err)
		return items
	}
	return items
}
