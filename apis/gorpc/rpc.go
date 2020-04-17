package gorpc

import (
	"go1234.cn/newResk/infra"
	"go1234.cn/newResk/infra/base"
)

type GoRPCApiStarter struct {
	infra.BaseStarter
}

//在流程中的Init阶段来注册
func (g *GoRPCApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRpc))
}
