package apm

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GrpcClient 结构体封装了 *grpc.ClientConn，以便进行 gRPC 客户端操作
type GrpcClient struct {
	*grpc.ClientConn
}

// NewGrpcClient 创建并返回一个新的 gRPC 客户端，连接到指定的地址
// 参数 addr: 目标 gRPC 服务的地址
func NewGrpcClient(addr string) (*GrpcClient, error) {
	// 使用 grpc.Dial 建立连接，配置了无安全传输和拦截器
	conn, err := grpc.Dial(addr,
		grpc.WithInsecure(),                                      // 禁用 TLS 加密（不推荐用于生产环境）
		grpc.WithUnaryInterceptor(unaryInterceptor()),            // 使用自定义的拦截器
		grpc.WithTransportCredentials(insecure.NewCredentials()), // 进一步避免传输中使用证书
	)
	if err != nil {
		// 如果连接失败，触发 panic
		panic(err)
	}
	// 返回一个新的 GrpcClient 实例，包装了连接
	return &GrpcClient{conn}, nil
}

// unaryInterceptor 定义了一个简单的 gRPC 拦截器，用于每个客户端调用
// 这里的拦截器只是简单调用 invoker（即执行实际请求），可以在这里做其他操作，例如日志记录、认证等
func unaryInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		// 执行实际的 gRPC 调用
		err := invoker(ctx, method, req, reply, cc, opts...)
		// 目前没有对错误进行处理，直接返回错误
		return err
	}
}
