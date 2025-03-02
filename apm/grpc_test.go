package apm

import (
	"context"
	"github.com/0xweb-3/amp_demo/protos"
	"testing"
)

type helloSvc struct {
	protos.UnimplementedHelloServiceServer
}

func (h *helloSvc) Receive(ctx context.Context, hello *protos.Hello) (*protos.Hello, error) {
	return hello, nil
}

func Test_Grpc(t *testing.T) {
	go func() {
		s := NewGrpcServer(":8080")
		protos.RegisterHelloServiceServer(s, &helloSvc{})
		s.Start()
	}()

	// 启动客户端
	client, _ := NewGrpcClient("127.0.0.1:8080")
	res, err := protos.NewHelloServiceClient(client).Receive(context.Background(), &protos.Hello{
		Msg: "测试",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log(res)
}
