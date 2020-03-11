package newResk

import (
	"go1234.cn/newResk/infra"
	"go1234.cn/newResk/infra/base"
)

func init() {
	infra.Register(&base.PropsStarter{})
}
