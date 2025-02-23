package apm

import (
	"fmt"
	"net/http"
	"testing"
	"time"
)

// 测试 HttpServer 启动和路由注册
func TestHttpServer(t *testing.T) {
	// 创建一个新的 HttpServer 实例
	server := NewHttpSever(":8080")

	// 注册一个简单的路由
	server.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	// 启动服务器
	go server.Start()

	// 等待服务器启动
	time.Sleep(time.Second * 5)

	//// 使用 httptest 请求服务器
	//req, err := http.NewRequest("GET", "http://localhost:8080/hello", nil)
	//if err != nil {
	//	t.Fatal(err)
	//}
	//
	//// 创建一个响应记录器
	//rr := httptest.NewRecorder()
	//
	//// 调用路由的处理函数
	//server.mux.ServeHTTP(rr, req)
	//
	//// 检查响应状态码
	//if status := rr.Code; status != http.StatusOK {
	//	t.Errorf("expected status code %v, got %v", http.StatusOK, status)
	//}
	//
	//// 检查响应内容
	//expected := "Hello, World!"
	//if rr.Body.String() != expected {
	//	t.Errorf("expected body %v, got %v", expected, rr.Body.String())
	//}
	//
	//// 优雅关闭服务器
	//err = server.close()
	//if err != nil {
	//	t.Fatalf("failed to close server: %v", err)
	//}
}

// 测试 HttpServer 启动失败的情况（例如端口被占用）
func TestHttpServer_StartFail(t *testing.T) {
	// 创建一个新的 HttpServer 实例，并设置端口为 8080
	server := NewHttpSever(":8080")

	// 启动第一个服务器
	go server.Start()

	// 等待服务器启动
	time.Sleep(time.Second)

	// 创建第二个服务器实例，尝试绑定到同一端口
	server2 := NewHttpSever(":8080")
	server2.Start()
}
