package model

import "github.com/0xweb-3/amp_demo/common/db"

type Order struct {
	db.BaseModel
	//OrderId string `gorm:"type:varchar(200);not null;default:''"`
	SkuId uint64 `gorm:"type:int(11);not null;default:0"`
	Num   int    `gorm:"type:int(11);not null;default:0"`
	Price int    `gorm:"type:int(11);not null;default:0"`
	Uid   uint64 `gorm:"type:int(11);not null;default:0"`
}
