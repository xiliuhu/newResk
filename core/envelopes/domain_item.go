package envelopes

import (
	"context"
	"github.com/segmentio/ksuid"
	"github.com/tietang/dbx"
	"github.com/ztaoing/infra/base"
	"github.com/ztaoing/newResk/services"
)

//抢红包的业务领域层
type itemDomain struct {
	RedEnvelopeItem
}

//生成itemNum
func (i *itemDomain) createItemNum() {
	i.ItemNum = ksuid.New().Next().String()
}

//创建item
func (i *itemDomain) Create(item services.RedEnvelopeItemDTO) {
	i.RedEnvelopeItem.FromDTO(&item)
	i.RecvUsername.Valid = true
	i.createItemNum()
}

//保存item
func (i *itemDomain) Save(ctx context.Context) (id int64, err error) {
	err = base.ExecuteContext(ctx, func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeItemDao{runner: runner}
		id, err = dao.Insert(&i.RedEnvelopeItem)
		return err
	})
	return id, err
}

//通过itemNum查询抢红包明细
func (i *itemDomain) GetOne(ctx context.Context, itemNum string) (dto *services.RedEnvelopeItemDTO) {
	err := base.ExecuteContext(ctx, func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeItemDao{runner: runner}
		po := dao.GetOne(itemNum)
		if po != nil {
			dto = po.ToDTO()
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return dto
}

//通过envelopeNum查询已抢到的红包列表

func (i *itemDomain) FindItems(envelopeNum string) (itemDTOs []*services.RedEnvelopeItemDTO) {
	var items []*RedEnvelopeItem
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeItemDao{runner: runner}
		items = dao.FindItems(envelopeNum)
		return nil
	})
	if err != nil {
		return itemDTOs
	}
	itemDTOs = make([]*services.RedEnvelopeItemDTO, 0)
	//转换
	for _, po := range items {
		itemDTOs = append(itemDTOs, po.ToDTO())
	}
	return itemDTOs
}

func (i *itemDomain) GetByUser(userId, envelopeNo string) (dto *services.RedEnvelopeItemDTO) {
	err := base.Tx(func(runner *dbx.TxRunner) error {
		dao := RedEnvelopeItemDao{runner: runner}
		po := dao.GetByUser(envelopeNo, userId)
		if po != nil {
			dto = po.ToDTO()
		}
		return nil
	})
	if err != nil {
		return nil
	}
	return dto
}
