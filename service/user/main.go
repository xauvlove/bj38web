package main

import (
	"bj38web/service/user/handler"
	user "bj38web/service/user/proto"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.user"),
		micro.Version("latest"),
	)
	user.RegisterUserHandler(service.Server(), new(handler.User))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
