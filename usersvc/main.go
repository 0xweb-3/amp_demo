package usersvc

import "github.com/0xweb-3/amp_demo/common/db"

type Sku struct {
	db.BaseModel
	Name  string `gorm:"type:varchar(20);not null;default:''"`
	Price int    `gorm:"type:int(11);not null;default:0"`
	Num   int    `gorm:"type:int(11);not null;default:0"`
}

type Order struct {
	db.BaseModel
	OrderId string `gorm:"type:varchar(200);not null;default:''"`
	SkuId   string `gorm:"type:varchar(200);not null;default:''"`
	Num     int    `gorm:"type:int(11);not null;default:0"`
	Price   int    `gorm:"type:int(11);not null;default:0"`
	Uid     uint64 `gorm:"type:int(11);not null;default:0"`
}
