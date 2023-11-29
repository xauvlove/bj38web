package controller

import (
	"bj38web/web/proto/getCaptcha"
	"bj38web/web/utils"
	"context"
	"encoding/json"
	"fmt"
	"image/png"
	"net/http"

	"github.com/afocus/captcha"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

func GetSession(c *gin.Context) {
	resp := make(map[string]string)
	resp["errno"] = utils.SYSTEM_ERROR
	resp["errmsg"] = utils.RecodeText(utils.SYSTEM_ERROR)
	c.JSON(http.StatusOK, resp)
}

func GetImageCd(c *gin.Context) {
	// 指定服务发现
	registry := consul.NewRegistry()
	consulClient := micro.NewService(micro.Registry(registry))
	// 初始化客户端
	microClient := getCaptcha.NewGetCaptchaService("go.micro.srv.getCaptcha", consulClient.Client())
	// 调用服务端接口
	uuid := c.Param("uuid")
	response, err := microClient.Call(context.TODO(), &getCaptcha.Request{Uuid: uuid})
	if err != nil {
		fmt.Printf("%v", err)
	}
	// 反序列化字节流，变为 img
	var img captcha.Image
	err = json.Unmarshal(response.B, &img)
	// 写浏览器数据
	png.Encode(c.Writer, img)
}
