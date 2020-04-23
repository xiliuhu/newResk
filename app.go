package newResk

import (
	_ "github.com/ztaoing/account/core/accounts"
	"github.com/ztaoing/infra"
	"github.com/ztaoing/infra/base"
	"go1234.cn/newResk/apis/gorpc"
	_ "go1234.cn/newResk/apis/gorpc"
	_ "go1234.cn/newResk/apis/web"
	_ "go1234.cn/newResk/core/envelopes"
	"go1234.cn/newResk/jobs"
)

func init() {

	infra.Register(&base.PropsStarter{})
	//infra.Register(&base.DbxStarter{})
	infra.Register(&base.ValidatorStart{})
	infra.Register(&base.GoRPCStarter{})
	infra.Register(&gorpc.GoRPCApiStarter{})
	infra.Register(&jobs.RefundExpiredStarter{})
	infra.Register(&base.IrisSveverStarter{})
	infra.Register(&infra.WebApiStart{})
	infra.Register(&base.EurekaStarter{})
	//infra.Register(&base.HookStarter{})
}
