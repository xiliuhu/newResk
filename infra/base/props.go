package base

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"go1234.cn/newResk/infra"
)

var props kvs.ConfigSource

//供外部调用
func Props() kvs.ConfigSource {
	return props
}

type PropsStarter struct {
	infra.BaseStarter
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {

	props = ctx.Props()
	fmt.Println("初始化配置")

}
