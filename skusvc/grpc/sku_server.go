package grpc

import (
	"context"
	"github.com/0xweb-3/amp_demo/protos"
	"github.com/0xweb-3/amp_demo/skusvc/dao"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SkuServer struct {
	protos.UnimplementedSkuServiceServer
}

func (s *SkuServer) decreaseStock(ctx context.Context, sku *protos.Sku) (*protos.Sku, error) {
	// 获取商品信息
	info, _ := dao.SkuDao.Get(ctx, sku.GetId())
	if info == nil {
		return nil, status.Error(codes.NotFound, "sku not found")
	}

	//  进行扣减库存
	err := dao.SkuDao.Decr(ctx, sku.GetId(), info.GetNum())
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Error())
	}
	return info, nil
}
