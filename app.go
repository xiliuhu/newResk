package newResk

import (
	"go1234.cn/newResk/infra"
	"go1234.cn/newResk/infra/base"
)

func init() {
	//
	infra.Register(&base.PropsStarter{})
	//dbx
	//infra.Register(&base.DbxStarter{})
	infra.Register(&base.ValidatorStart{})
	infra.Register(&base.IrisSveverStarter{})
}
