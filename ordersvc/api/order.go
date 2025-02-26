package api

import (
	"context"
	"github.com/0xweb-3/amp_demo/apm"
	"github.com/0xweb-3/amp_demo/ordersvc/grpcclient"
	"github.com/0xweb-3/amp_demo/ordersvc/model"
	"github.com/0xweb-3/amp_demo/protos"
	"net/http"
	"strconv"
)

type order struct{}

var Order = &order{}

func (o *order) Add(w http.ResponseWriter, r *http.Request) {
	// 获取参数
	values := r.URL.Query()
	var (
		uid, _   = strconv.Atoi(values.Get("uid"))
		skuId, _ = strconv.Atoi(values.Get("sku_id"))
		num, _   = strconv.Atoi(values.Get("num"))
	)
	// todo 检查用户信息

	// todo 库存的扣减
	_, err := grpcclient.SkuClient.DecreaseStock(context.TODO(), &protos.Sku{
		Id:  10,
		Num: 1,
	})
	if err != nil {
		apm.Logger.Error(context.TODO(), "Order_Add", map[string]any{
			"uid":   uid,
			"skuId": skuId,
			"num":   num,
		}, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//创建订单
	err = apm.Infra.DB.Save(&model.Order{
		SkuId: uint64(skuId),
		Num:   num,
		Price: 100,
		Uid:   uint64(uid),
	}).Error
	if err != nil {
		apm.Logger.Error(context.TODO(), "Order_Add", map[string]any{
			"uid":   uid,
			"skuId": skuId,
			"num":   num,
		}, err)

		apm.HttpStatus.Error(w, err.Error(), nil)
	}
	apm.HttpStatus.OK(w)
}

func (o *order) Get(w http.ResponseWriter, r *http.Request) {
	var firstOrder model.Order
	err := apm.Infra.DB.First(&firstOrder).Error
	if err != nil {
		apm.Logger.Error(context.TODO(), "Order_Add", map[string]any{}, err)
		apm.HttpStatus.Error(w, err.Error(), nil)
	}

	apm.HttpStatus.OKBody(w, "", firstOrder)
}
