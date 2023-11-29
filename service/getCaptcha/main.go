package main

import (
	"bj38web/service/getCaptcha/handler"
	"bj38web/service/getCaptcha/proto/getCaptcha"
	"fmt"

	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// 初始化 consul
	registry := consul.NewRegistry()

	// new service
	service := micro.NewService(
		micro.Name("go.micro.srv.getCaptcha"),
		micro.Version("latest"),
		// 防止随机生成 port
		micro.Address("localhost:13363"),
		// 注册中心
		micro.Registry(registry),
	)
	// register service
	getCaptcha.RegisterGetCaptchaHandler(service.Server(), new(handler.GetCaptcha))
	// run service
	if err := service.Run(); err != nil {
		fmt.Printf("run service error = %v\n", err)
	}

}
