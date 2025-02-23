package apm

import (
	"testing"
)

func Test_infra_Init(t *testing.T) {
	Infra.Init(
		InfraDbOption("root:yjfc4883212@tcp(192.168.21.2:3315)/amp?charset=utf8mb4&parseTime=True&loc=Local"),
		InfraRDbOption("192.168.21.2", 6385),
	)

	t.Log(Infra)
}
