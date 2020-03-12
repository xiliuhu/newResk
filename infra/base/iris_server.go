package base

var irisApplication *iris.Application

func Iris() *iris.Application {
	return irisApplication
}

type IrisSveverStarter struct {
	infra.BaseStarter
}

//iris需要在 Init和Start阶段处理
func (i *IrisSveverStarter) Init() {
	//在此阶段需要做以下事情
	//创建iris application实例
	//日志组件配置和扩展
	//主要中间件的配置：recover，日志输出中间件的自定义

}

func (i *IrisSveverStarter) Start() {

}
