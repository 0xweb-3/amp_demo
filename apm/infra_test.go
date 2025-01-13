package apm

import (
	"testing"
)

func Test_infra_Init(t *testing.T) {
	Infra.Init(
		InfraDbOption("root:yjfc4883212@tcp(127.0.0.1:3315)/amp?charset=utf8mb4&parseTime=True&loc=Local"),
		InfraRDbOption("127.0.0.1", 6385),
	)
}
