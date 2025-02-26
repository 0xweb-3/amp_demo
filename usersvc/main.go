package main

import (
	"github.com/0xweb-3/amp_demo/apm"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// 初始化db,http，grpcClient
	dsn := "root:yjfc4883212@tcp(192.168.21.2:3315)/amp?charset=utf8mb4&parseTime=True&loc=Local"
	apm.Infra.Init(apm.InfraDbOption(dsn))

	// todo grpcClient 初始化
	ginServer := apm.NewGinServer(":8082")

	// 注入处理函数
	ginServer.Handle(http.MethodGet, "/test", func(ctx *gin.Context) {
		apm.GinStatus.OK(ctx)
	})

	//.HandleFunc("/order/add", api.Order.Add)
	//ginServer.HandleFunc("/order/get", api.Order.Get)

	// 启动所有封装的服务
	apm.EndPoint.Start()
}
