package main

import (
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	_ "go1234.cn/newResk"
	"go1234.cn/newResk/infra"
)

func main() {
	filePath := kvs.GetCurrentFilePath("config.ini", 1)
	//加载配置文件
	conf := ini.NewIniFileCompositeConfigSource(filePath)
	app := infra.New(conf)
	app.Start()
}
