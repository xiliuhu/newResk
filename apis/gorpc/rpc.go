package gorpc

import (
	"github.com/ztaoing/infra"
	"github.com/ztaoing/infra/base"
)

type GoRPCApiStarter struct {
	infra.BaseStarter
}

//在流程中的Init阶段来注册
func (g *GoRPCApiStarter) Init(ctx infra.StarterContext) {
	base.RpcRegister(new(EnvelopeRpc))
}
