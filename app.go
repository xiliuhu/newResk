package newResk

import (
	_ "github.com/ztaoing/account/core/accounts"
	"github.com/ztaoing/infra"
	"github.com/ztaoing/infra/base"
	"github.com/ztaoing/newResk/apis/gorpc"
	_ "github.com/ztaoing/newResk/apis/gorpc"
	_ "github.com/ztaoing/newResk/apis/web"
	_ "github.com/ztaoing/newResk/core/envelopes"
	"github.com/ztaoing/newResk/jobs"
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
