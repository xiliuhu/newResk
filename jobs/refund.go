package jobs

import (
	"github.com/sirupsen/logrus"
	"go1234.cn/newResk/infra"
	"time"
)

//退款
type RefundExpiredStarter struct {
	infra.BaseStarter
	ticker *time.Ticker
}

func (r *RefundExpiredStarter) Init(ctx infra.StarterContext) {
	//时间间隔
	d := ctx.Props().GetDurationDefault("jobs.refund.interval", time.Minute)
	r.ticker = time.NewTicker(d)
}

func (r *RefundExpiredStarter) Start(ctx infra.StarterContext) {
	//使用goroutine来执行定时
	go func() {
		for {
			c := <-r.ticker.C
			logrus.Debug("触发红包退款...", c)
			//红包退款

		}
	}()
}

//停止的时候，定时任务也需要停止
func (r *RefundExpiredStarter) Stop(ctx infra.StarterContext) {
	//停止定时触发红包退款
	r.ticker.Stop()
}
