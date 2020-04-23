package main

import (
	"github.com/tietang/go-eureka-client/eureka"
	"time"
)

func main() {
	config := eureka.Config{
		DialTimeout: time.Second * 10,
	}
	client := eureka.NewClientByConfig([]string{
		"http://127.0.0.1:8761/eureka",
	}, config)
	appName := "GO-Example"
	//设置实例
	instance := eureka.NewInstanceInfo("test.com", appName, "127.0.0.2", 8080, 30, false)
	//注册实例
	client.RegisterInstance(appName, instance)
	client.Start()

}
