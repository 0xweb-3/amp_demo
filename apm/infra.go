package apm

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// 基础设施
type infra struct {
	DB  *gorm.DB
	Rdb *redis.Client
}

var Infra = &infra{}

type InfraOption func(*infra)

// "root:xinbingliang@tcp(127.0.0.1:3310)/fishline?charset=utf8mb4&parseTime=True&loc=Local"
func InfraDbOption(connectUrl string) InfraOption {
	return func(infra *infra) {
		var err error
		infra.DB, err = gorm.Open(mysql.Open(connectUrl), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, //去除表明后的s
			},
		})
		if err != nil {
			panic(err)
		}
	}
}

func InfraRDbOption(host string, port int) InfraOption {
	return func(infra *infra) {
		var err error
		addr := fmt.Sprintf("%s:%d", host, port)
		Infra.Rdb = redis.NewClient(&redis.Options{
			Addr: addr,
			DB:   1,
		})
		// 检查链接是否成功
		ctx := context.Background()
		_, err = Infra.Rdb.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
	}
}

func (i *infra) Init(options ...InfraOption) {
	for _, option := range options {
		option(i)
	}
}
