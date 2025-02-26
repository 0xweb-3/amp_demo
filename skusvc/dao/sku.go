package dao

import (
	"context"
	"github.com/0xweb-3/amp_demo/apm"
	"github.com/0xweb-3/amp_demo/protos"
	"github.com/0xweb-3/amp_demo/skusvc/dao/model"
	"gorm.io/gorm"
)

type skuDao struct {
}

var SkuDao = &skuDao{}

func (s *skuDao) Get(ctx context.Context, skuId uint64) (*protos.Sku, error) {
	sku := &model.Sku{}
	err := apm.Infra.DB.First(sku, skuId).Error
	if err != nil {
		return nil, err
	}
	return &protos.Sku{
		Name:  sku.Name,
		Id:    sku.ID,
		Price: sku.Price,
		Num:   sku.Num,
	}, nil
}

func (s *skuDao) Decr(ctx context.Context, id uint64, num int32) error {
	// 查找 SKU 并直接扣减库存
	err := apm.Infra.DB.Model(&model.Sku{}).
		Where("id = ?", id).
		UpdateColumn("num", gorm.Expr("num - ?", num)).Error

	return err
}
