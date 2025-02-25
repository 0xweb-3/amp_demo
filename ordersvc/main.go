package main

import (
	"github.com/0xweb-3/amp_demo/apm"
	"github.com/0xweb-3/amp_demo/ordersvc/api"
	"net/http"
)

func main() {
	// 初始化db,http，grpcClient
	dsn := "root:yjfc4883212@tcp(192.168.21.2:3315)/amp?charset=utf8mb4&parseTime=True&loc=Local"
	apm.Infra.Init(apm.InfraDbOption(dsn))

	// todo grpcClient 初始化
	httpServer := apm.NewHttpSever(":8081")

	// 注入处理函数
	httpServer.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	httpServer.HandleFunc("/order/add", api.Order.Add)

	httpServer.HandleFunc("/order/get", api.Order.Get)

	// 启动所有封装的服务
	apm.EndPoint.Start()
}
