package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/tietang/props/consul"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"github.com/ztaoing/infra"
	"github.com/ztaoing/infra/base"
)

//从consul中读取配置
func main() {
	//获取配置文件所在路径
	file := kvs.GetCurrentFilePath("boot.ini", 1)
	//加载和解析配置文件
	conf := ini.NewIniFileCompositeConfigSource(file)
	//通过consul的配置信息来读取consul
	address := conf.GetDefault("consul.address", "127.0.0.1:8500")
	contexts := conf.KeyValue("consul.contexts").Strings()
	consulConf := consul.NewCompositeConsulConfigSourceByType(contexts, address, kvs.ContentIni)
	//合并配置
	consulConf.Add(conf)
	base.InitLog(consulConf)
	app := infra.New(conf)
	app.Start()

}
