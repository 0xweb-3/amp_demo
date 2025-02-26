package main

import (
	"github.com/0xweb-3/amp_demo/apm"
	"github.com/0xweb-3/amp_demo/protos"
	"github.com/0xweb-3/amp_demo/skusvc/grpc"
)

func main() {
	// 初始化db,http，grpcClient
	dsn := "root:yjfc4883212@tcp(192.168.21.2:3315)/amp?charset=utf8mb4&parseTime=True&loc=Local"
	apm.Infra.Init(apm.InfraDbOption(dsn))

	grpcServer := apm.NewGrpcServer(":5001")

	//注入gpc实现
	protos.RegisterSkuServiceServer(grpcServer, &grpc.SkuServer{})

	// 启动所有封装的服务
	apm.EndPoint.Start()
}
