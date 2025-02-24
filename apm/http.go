package apm

import (
	"context"
	"net/http"
)

// HttpServer 自定义的 HTTP 服务器类型，它组合了 http.ServeMux 和 http.Server。
// ServeMux 是 Go 的默认多路复用器（路由器），Server 是标准的 HTTP 服务器。
// HttpServer 结构体封装了常见的 HTTP 服务器功能，并为路由注册和服务器启动提供了更简便的接口。
type HttpServer struct {
	mux          *http.ServeMux // 路由器，用于将 HTTP 请求分发到不同的处理函数
	*http.Server                // 内嵌 http.Server，提供基础的 HTTP 服务器功能
}

// NewHttpServer 是 HttpServer 的构造函数，创建一个新的 HttpServer 实例。
// 它初始化了一个 ServeMux 和一个 http.Server 实例，并将二者结合。
// addr 是服务器监听的地址，通常是 "localhost:8080" 这样的格式。
func NewHttpSever(addr string) *HttpServer {
	// 创建 ServeMux 路由器
	mux := http.NewServeMux()
	// 创建 http.Server 实例，指定监听的地址和使用 mux 作为请求处理器
	server := &http.Server{Addr: addr, Handler: mux}
	// 返回一个包含 mux 和 server 的 HttpServer 实例
	return &HttpServer{mux: mux, Server: server}
}

// Handle 是 HttpServer 的一个方法，用于注册处理指定 pattern 的请求。
// pattern 是 URL 路径模式，handler 是处理请求的具体函数。
func (h *HttpServer) Handle(pattern string, handler http.Handler) {
	h.mux.Handle(pattern, handler) // 将请求与 handler 关联
}

// HandleFunc 是 HttpServer 的一个方法，简化了 Handle 的使用，直接传入一个函数。
// handler 是一个函数，接收 http.ResponseWriter 和 http.Request 参数，处理请求并返回响应。
func (h *HttpServer) HandleFunc(pattern string, handler func(w http.ResponseWriter, r *http.Request)) {
	h.mux.HandleFunc(pattern, handler) // 将请求与 handler 函数关联
}

// Start 启动 HTTP 服务器，它在一个 goroutine 中运行 ListenAndServe。
// 通过 goroutine 启动服务器，使得主线程不会阻塞，能够继续执行其他操作。
// 如果服务器启动失败，会触发 panic。
func (h *HttpServer) Start() {
	go func() {
		err := h.ListenAndServe() // 启动 HTTP 服务器
		if err != nil {
			panic(err) // 如果遇到错误（比如端口被占用），抛出 panic
		}
	}()
}

// close 优雅关闭 HTTP 服务器，调用 Shutdown 方法来停止服务器的运行。
// 通过 context.Background() 确保即使在没有超时的情况下，也能平滑关闭。
// 该方法返回服务器关闭时的错误信息。
func (h *HttpServer) close() error {
	return h.Shutdown(context.Background()) // 调用 Shutdown 方法停止 HTTP 服务器
}

//func helloHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Fprintf(w, "Hello, World!")
//}
//
//func main() {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/hello.proto", helloHandler)
//
//	server := &http.Server{
//		Addr:    ":8080",
//		Handler: mux,
//	}
//
//	fmt.Println("Starting server at :8080")
//	if err := server.ListenAndServe(); err != nil {
//		fmt.Println(err)
//	}
//}
