package api

import (
	"github.com/0xweb-3/amp_demo/apm"
	"github.com/gin-gonic/gin"
)

func Add(ctx *gin.Context) error {
	// todo 补充代码
	apm.GinStatus.OKBody(ctx, "", map[string]interface{}{
		"id": 100,
	})
}
