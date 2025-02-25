package apm

import (
	"os"
	"os/signal"
	"syscall"
)

type starter interface {
	Start()
}

type closer interface {
	Close()
}

var (
	globalStarters = make([]starter, 0)
	globalClosers  = make([]closer, 0)
)

type endPoint struct {
	stop chan struct{}
}

var EndPoint = &endPoint{make(chan struct{}, 1)}

func (e *endPoint) Start() {
	for _, com := range globalStarters {
		com.Start()
	}
	go func() {
		// 监听服务的结束
		quit := make(chan os.Signal)
		signal.Notify(quit, syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM)
		<-quit
		e.Shutdown()
	}()
	<-e.stop
}

func (e *endPoint) Shutdown() {
	for _, com := range globalClosers {
		com.Close()
	}
	e.stop <- struct{}{}
}
