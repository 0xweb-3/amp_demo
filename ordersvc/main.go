package main

import (
	"github.com/0xweb-3/amp_demo/apm"
	"github.com/0xweb-3/amp_demo/ordersvc/api"
	"github.com/0xweb-3/amp_demo/ordersvc/grpcclient"
	"github.com/0xweb-3/amp_demo/protos"
	"net/http"
)

func main() {
	// 初始化db,http，grpcClient
	dsn := "root:yjfc4883212@tcp(192.168.21.2:3315)/amp?charset=utf8mb4&parseTime=True&loc=Local"
	otelEndpoint := "127.0.0.1:54317"
	apm.Infra.Init(
		apm.InfraDbOption(dsn),
		apm.InfraEnableApm(otelEndpoint),
	)

	// 对sku服务客户端进行初始化
	skuClient, err := apm.NewGrpcClient(":5001")
	if err != nil {
		panic(err)
	}
	grpcclient.SkuClient = protos.NewSkuServiceClient(skuClient)

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
