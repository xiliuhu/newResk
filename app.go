package newResk

import (
	_ "go1234.cn/newResk/apis/web"
	_ "go1234.cn/newResk/core/accounts"
	"go1234.cn/newResk/infra"
	"go1234.cn/newResk/infra/base"
)

func init() {

	infra.Register(&base.PropsStarter{})
	//infra.Register(&base.DbxStarter{})
	infra.Register(&base.ValidatorStart{})
	infra.Register(&base.IrisSveverStarter{})
	infra.Register(&infra.WebApiStart{})
}
