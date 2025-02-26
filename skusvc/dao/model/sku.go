package model

import "github.com/0xweb-3/amp_demo/common/db"

type Sku struct {
	db.BaseModel
	Name  string `gorm:"type:varchar(20);not null;default:''"`
	Price int    `gorm:"type:int(11);not null;default:0"`
	Num   int    `gorm:"type:int(11);not null;default:0"`
}
