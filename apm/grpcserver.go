package apm

import (
	"context"
	"google.golang.org/grpc"
	"net"
)

// GrpcServer 结构体封装了 *grpc.Server 和地址信息
// addr: 服务器监听的地址
type GrpcServer struct {
	*grpc.Server        // 继承 gRPC 服务器的功能
	addr         string // 服务器地址
}

func (g *GrpcServer) Close() {
	g.Server.GracefulStop()
}

// NewGrpcServer 创建并返回一个新的 gRPC 服务器，指定了监听地址
// 参数 addr: gRPC 服务监听的地址
func NewGrpcServer(addr string) *GrpcServer {
	// 创建一个新的 gRPC 服务器，并配置了拦截器
	svc := grpc.NewServer(grpc.UnaryInterceptor(unaryServerInterceptor())) // 使用自定义的拦截器
	//return &GrpcServer{svc, addr}                                          // 返回一个封装了 gRPC 服务器的 GrpcServer 实例
	server := &GrpcServer{svc, addr}
	globalStarters = append(globalStarters, server)
	globalClosers = append(globalClosers, server)
	return server
}

// start 启动 gRPC 服务器并开始监听指定的地址
// 这个方法在 goroutine 中异步执行 gRPC 服务器的服务
func (g *GrpcServer) Start() {
	// 创建 TCP 监听器，绑定到 g.Server.addr
	l, err := net.Listen("tcp", g.addr)
	if err != nil {
		// 监听失败时触发 panic
		panic(err)
	}
	// 使用 goroutine 启动 gRPC 服务器，并让它在后台运行
	go func() {
		// 启动 gRPC 服务器，开始处理传入的请求
		err = g.Server.Serve(l)
		if err != nil {
			// 如果启动失败或服务器停止时触发 panic
			panic(err)
		}
	}()
}

// unaryServerInterceptor 定义了一个 gRPC 服务器拦截器，拦截每个 unary 请求
// 在请求处理之前和之后可以执行自定义逻辑（目前只是调用 handler 继续处理请求）
func unaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		// 调用实际的处理函数
		return handler(ctx, req)
	}
}
