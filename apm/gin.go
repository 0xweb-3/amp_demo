package apm

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GinHttpServer 自定义的 Gin HTTP 服务器类型
type GinHttpServer struct {
	engine *gin.Engine
	server *http.Server
}

// NewGinServer 创建一个新的 GinHttpServer 实例
func NewGinServer(addr string) *GinHttpServer {
	// 创建一个 Gin 引擎实例
	engine := gin.Default()
	// 创建一个 http.Server 实例
	server := &http.Server{
		Addr:    addr,
		Handler: engine,
	}
	// 返回 GinHttpServer 实例
	s := &GinHttpServer{engine: engine, server: server}
	globalStarters = append(globalStarters, s)
	globalClosers = append(globalClosers, s)
	return s
}

// Handle 注册一个路由和对应的处理函数
func (g *GinHttpServer) Handle(method, pattern string, handler gin.HandlerFunc) {
	switch method {
	case http.MethodGet:
		g.engine.GET(pattern, handler)
	case http.MethodPost:
		g.engine.POST(pattern, handler)
	case http.MethodPut:
		g.engine.PUT(pattern, handler)
	case http.MethodDelete:
		g.engine.DELETE(pattern, handler)
	default:
		g.engine.Any(pattern, handler)
	}
}

// Start 启动 HTTP 服务器
func (g *GinHttpServer) Start() {
	go func() {
		if err := g.server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()
}

// Close 优雅关闭服务器
func (g *GinHttpServer) Close() {
	g.server.Shutdown(context.Background())
}
