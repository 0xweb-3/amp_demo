package model

import (
	"github.com/0xweb-3/amp_demo/common/db"
)

type User struct {
	db.BaseModel
	Name string `gorm:"type:varchar(20);not null;default:''"`
}
