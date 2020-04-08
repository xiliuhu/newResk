package accounts

import (
	"github.com/go-playground/validator/v10"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
	"go1234.cn/newResk/infra/base"
	"go1234.cn/newResk/services"
	"time"
)

//应用服务层
var _ services.AccountService = new(accountService)

type accountService struct {
}

func (a *accountService) CreateAccount(dto services.AccountCreateDTO) (*services.AccountDTO, error) {
	domain := accountDomain{}
	//验证输入参数
	err := base.Validate().Struct(&dto)
	if err != nil {
		_, ok := err.(*validator.InvalidValidationError)
		if ok {
			logrus.Error("验证错误", err)
		}
		errs, ok := err.(*validator.ValidationErrors)
		if ok {
			for _, value := range errs {
				logrus.Error(value.Translate(base.Translate()))
			}
		}
		return nil, err
	}
	amount, err := decimal.NewFromString(dto.Amount)
	//执行账户创建的业务逻辑
	account := services.AccountDTO{
		AccountName:  dto.AccountName,
		AccountType:  dto.AccountType,
		CurrencyCode: dto.CurrencyCode,
		UserId:       dto.UserId,
		Username:     dto.Username,
		Balance:      amount,
		Status:       1,
	}
	accountDto, err := domain.CreateAccount(account)
	return accountDto, err
}

func (a *accountService) Transfer(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	panic("implement me")
}

func (a *accountService) StoreValue(dto services.AccountTransferDTO) (services.TransferedStatus, error) {
	panic("implement me")
}

func (a *accountService) GetEnvelopeAccountByUserId(userId string) *services.AccountDTO {
	panic("implement me")
}
