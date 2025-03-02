package apm

import (
	"context"
	"testing"
)

func Test_infra_Init(t *testing.T) {
	Infra.Init(
		InfraDbOption("root:yjfc4883212@tcp(192.168.21.2:3315)/amp?charset=utf8mb4&parseTime=True&loc=Local"),
		InfraRDbOption("192.168.21.2", 6385),
	)

	t.Log(Infra)
}

func TestInfraRDbOption(t *testing.T) {
	Infra.Init(
		InfraRDbOption("192.168.21.2", 6385),
		InfraEnableApm("127.0.0.1:54317"),
	)

	_, _ = Infra.Rdb.Get(context.TODO(), "xxxx").Result()
	EndPoint.Shutdown()
}
