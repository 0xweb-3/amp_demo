package apm

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// ginStatus 用于 Gin 封装的结构体
type ginStatus struct {
}

var GinStatus = &ginStatus{}

// OK 返回成功的响应
func (h *ginStatus) OK(c *gin.Context) {
	status := &Status{Code: http.StatusOK, Message: "success"}

	c.JSON(http.StatusOK, status)
}

// OKBody 携带具体业务信息的成功响应
func (h *ginStatus) OKBody(c *gin.Context, msg string, body any) {
	status := &Status{Code: http.StatusOK, Message: msg, Body: body}

	c.JSON(http.StatusOK, status)
}

// Fail 业务级别错误的响应
func (h *ginStatus) Fail(c *gin.Context, msg string, body any) {
	status := &Status{Code: http.StatusBadRequest, Message: msg, Body: body}

	c.JSON(http.StatusBadRequest, status)
}

// Error 服务器级别错误的响应
func (h *ginStatus) Error(c *gin.Context, msg string, body any) {
	status := &Status{Code: http.StatusInternalServerError, Message: msg, Body: body}

	c.JSON(http.StatusInternalServerError, status)
}
