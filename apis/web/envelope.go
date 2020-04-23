package web

import (
	"github.com/kataras/iris"
	"github.com/ztaoing/infra"
	"github.com/ztaoing/infra/base"
	"go1234.cn/newResk/services"
)

func init() {
	//注册初始化函数
	infra.RegisterApi(&EnvelopeApi{})
}

type EnvelopeApi struct {
	//应用服务层
	service services.RedEnvelopeService
}

func (e *EnvelopeApi) Init() {
	e.service = services.GetRedEnvelopeService()
	//定义router
	groupRouter := base.Iris().Party("/v1/envelope")
	groupRouter.Post("/send", e.sendHandler)
	groupRouter.Post("/receive", e.receiveHandler)
}

//发红包
func (e *EnvelopeApi) sendHandler(ctx iris.Context) {
	//从请求中读取数据，存储到services.RedEnvelopeSendDTO{}中
	dto := services.RedEnvelopeSendDTO{}
	err := ctx.ReadJSON(&dto)
	res := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		res.Code = base.ResCodeRequestParamsError
		res.Message = err.Error()
		//通过json函数将信息写入到response的body中
		ctx.JSON(&res)
		return
	}
	//正常读取到了数据,就执行发送红包
	activity, err := e.service.Send(dto)
	if err != nil {
		res.Code = base.ResCodeInnerServerError
		res.Message = err.Error()
		//通过json函数将信息写入到response的body中
		ctx.JSON(&res)
		return
	}
	//发送成功返回红包链接
	res.Data = activity
	ctx.JSON(res)

}

//抢红包
func (e *EnvelopeApi) receiveHandler(ctx iris.Context) {
	dto := services.RedEnvelopeReceiveDTO{}
	err := ctx.ReadJSON(&dto)
	res := base.Res{
		Code: base.ResCodeOk,
	}
	if err != nil {
		res.Code = base.ResCodeRequestParamsError
		res.Message = err.Error()
		//通过json函数将信息写入到response的body中
		ctx.JSON(&res)
		return
	}
	//正常读取到了数据,就执行发送红包
	item, err := e.service.Recive(dto)
	if err != nil {
		res.Code = base.ResCodeInnerServerError
		res.Message = err.Error()
		//通过json函数将信息写入到response的body中
		ctx.JSON(&res)
		return
	}
	//发送成功返回红包链接
	res.Data = item
	ctx.JSON(res)
}
