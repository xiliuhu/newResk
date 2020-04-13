package web

import (
	"github.com/kataras/iris"
	"go1234.cn/newResk/infra"
	"go1234.cn/newResk/infra/base"
	"go1234.cn/newResk/services"
)

//web api是基于iris的
//对以每一个子业务，定义统一的前缀

//资金账户 的根路径定义为：/account
//版本号：/v1/account

func init() {
	infra.RegisterApi(&AccountApi{})
}

type AccountApi struct {
}

func (a *AccountApi) Init() {
	groupRouter := base.Iris().Party("/v1/account")
	groupRouter.Post("/create", createHaddler)
}

//创建账户的接口:/v1/account/create
func createHaddler(ctx iris.Context) {

	//获取请求的参数
	account := services.AccountCreateDTO{}
	err := ctx.ReadJSON(&account)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	//未出错，创建账户
	service := services.GetAccountService()
	dto, err := service.CreateAccount(account)
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
	}
	r.Data = dto
	ctx.JSON(&r)

}

//转账的接口: /v1/account/transfer
func transferHandler(ctx iris.Context) {
	//获取请求的参数
	account := services.AccountTransferDTO{}
	err := ctx.ReadJSON(&account)
	r := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
		ctx.JSON(&r)
		return
	}
	//未出错，执行转账
	service := services.GetAccountService()
	status, err := service.Transfer(account)
	if err != nil {
		r.Code = base.ResCodeRequestParamsError
		r.Message = err.Error()
	}
	r.Data = status
	//转账失败
	if status != services.TransferedStatusSuccess {
		r.Code = base.ResCodeBissTransferFailure
		r.Message = err.Error()

	}

	//转账成功
	ctx.JSON(&r)
}

//查询红包账户的接口: /v1/account/envelope/get

//查询账户信息的接口: /v1/account/get
