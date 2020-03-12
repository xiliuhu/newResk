package base

import (
	"github.com/sirupsen/logrus"
	"github.com/tietang/dbx"
	"github.com/tietang/props/kvs"
	"go1234.cn/newResk/infra"
)

//把数据库的初始化放在setup阶段
var database *dbx.database

func DbxDatabase() *dbx.database {
	return database
}

//dbx的starter
type DbxStarter struct {
	infra.BaseStarter
}

func (s *DbxStarter) Setup(ctx infra.StarterContext) {
	conf := ctx.Props()
	//初始化数据库配置
	settings := dbx.Settings{}
	err := kvs.Unmarshal(conf, settings, "mysql")
	if err != nil {
		panic(err)
	}
	logrus.Info("mysql.conn url:", settings.ShortDataSourceName)
	dbx, err := dbx.Open(settings)
	if err != nil {
		panic(err)
	}
	logrus.Info(dbx.Ping())
	database = dbx
}
