package apm

import (
	"context"
	"fmt"
	"github.com/0xweb-3/amp_demo/apm/internal"
	"github.com/redis/go-redis/v9"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.24.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"gorm.io/plugin/opentelemetry/tracing"
	"time"
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

		// 启用 OpenTelemetry 追踪
		if err = infra.DB.Use(tracing.NewPlugin()); err != nil {
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

		Infra.Rdb.AddHook(&redisHook{})
		// 检查链接是否成功
		ctx := context.Background()
		_, err = Infra.Rdb.Ping(ctx).Result()
		if err != nil {
			panic(err)
		}
	}
}

func InfraEnableApm(otelEndpoint string) InfraOption {
	return func(infra *infra) {
		ctx := context.Background()
		res, err := resource.New(ctx,
			resource.WithAttributes(
				semconv.ServiceName(internal.BuildInfo.AppName()),
			),
		)
		if err != nil {
			panic(err)
		}

		ctx, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		conn, err := grpc.DialContext(ctx, otelEndpoint,
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithBlock(),
			grpc.WithTimeout(3*time.Second), // 设置 gRPC 连接超时
		)
		if err != nil {
			panic(err)
		}

		traceExporter, err := otlptracegrpc.New(ctx, otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			panic(err)
		}

		bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
		tracerProvider := sdktrace.NewTracerProvider(
			sdktrace.WithSampler(sdktrace.AlwaysSample()),
			sdktrace.WithResource(res),
			sdktrace.WithSpanProcessor(bsp),
		)
		otel.SetTracerProvider(tracerProvider)
		otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{}))
		globalClosers = append(globalClosers, &traceProviderComponent{provider: tracerProvider})
	}
}

type traceProviderComponent struct {
	provider *sdktrace.TracerProvider
}

func (t *traceProviderComponent) Close() {
	t.provider.Shutdown(context.Background())
}

// Init 初始化方法
func (i *infra) Init(options ...InfraOption) {
	for _, option := range options {
		option(i)
	}
	Tracer = otel.Tracer(internal.BuildInfo.AppName())
}
