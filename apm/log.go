package apm

import (
	"context"
	"github.com/sirupsen/logrus"
)

func init() {
	// 设置 logrus 格式化输出为 JSON 格式
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

type log struct{}

// logger 是我们用来记录日志的全局实例
var logger = &log{}

// Info 方法用于记录 INFO 级别的日志
// ctx: 上下文对象，一般用于跟踪请求的生命周期
// action: 模块名称或操作名称
// kv: 额外的 key-value 数据，会添加到日志中
func (l *log) Info(ctx context.Context, action string, kv map[string]any) {
	kv["action"] = action        // 为日志添加 action 字段，标明日志来源
	logrus.WithFields(kv).Info() // 输出 Info 级别的日志
}

// Warn 方法用于记录 WARN 级别的日志
// ctx: 上下文对象，通常用于处理请求的生命周期
// action: 操作名称或模块名称
// kv: 日志的额外信息
func (l *log) Warn(ctx context.Context, action string, kv map[string]any) {
	kv["action"] = action        // 为日志添加 action 字段
	logrus.WithFields(kv).Warn() // 输出 Warn 级别的日志
}

// Debug 方法用于记录 Debug 级别的日志
// ctx: 上下文对象，用于关联请求或操作
// action: 模块名称或操作名称
// kv: 额外的 key-value 数据，用于日志记录
func (l *log) Debug(ctx context.Context, action string, kv map[string]any) {
	kv["action"] = action         // 为日志添加 action 字段
	logrus.WithFields(kv).Debug() // 输出 Debug 级别的日志
}

// Error 方法用于记录 Error 级别的日志
// ctx: 上下文对象，通常用于请求跟踪
// action: 操作名称或模块名称
// kv: 额外的日志信息
// err: 需要记录的错误信息
func (l *log) Error(ctx context.Context, action string, kv map[string]any, err error) {
	kv["action"] = action                        // 为日志添加 action 字段
	logrus.WithFields(kv).WithError(err).Error() // 输出 Error 级别的日志，并附带错误信息
}
