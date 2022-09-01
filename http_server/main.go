package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"

	"syscall"

	"github.com/ralvescosta/gokit/env"
	"github.com/ralvescosta/gokit/http/server"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit/otel/metric"
	"github.com/ralvescosta/gokit/otel/trace"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.uber.org/zap/zapcore"
)

func main() {
	ctx := context.Background()

	cfg, err := env.
		New().
		Tracing().
		HTTPServer().
		Build()
	if err != nil {
		panic(err)
	}

	logger, err := logging.NewDefaultLogger(cfg)
	if err != nil {
		panic(err)
	}

	closeTracerExporter, err := trace.
		NewOTLP(cfg, logger).
		WithApiKeyHeader().
		Build(ctx)
	if err != nil {
		panic(err)
	}
	defer closeTracerExporter(ctx)

	closeMetricExporter, err := metric.
		NewOTLP(cfg, logger).
		WithApiKeyHeader().
		Build(ctx)
	if err != nil {
		panic(err)
	}
	defer closeMetricExporter(ctx)

	meter := global.Meter("github.com/gokit_example/http_server")
	counter, err := meter.SyncFloat64().Counter("important_metric", instrument.WithDescription("Measures the cumulative /hello requests"))
	if err != nil {
		logger.Fatal("Failed to create the instrument", zapcore.Field{
			Key:       "err",
			Interface: err,
		})
	}
	counter.Add(ctx, 1.0)
	counter.Add(ctx, 1.0)
	counter.Add(ctx, 1.0)
	counter.Add(ctx, 1.0)
	counter.Add(ctx, 1.0)
	counter.Add(ctx, 1.0)
	counter.Add(ctx, 1.0)
	counter.Add(ctx, 1.0)

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	httpServer := server.
		New(cfg, logger, sig).
		WithTracing().
		Build()

	httpServer.RegisterRoute(http.MethodGet, "/hello", func(w http.ResponseWriter, r *http.Request) {
		counter.Add(ctx, 1.0)
		w.WriteHeader(200)
		w.Write([]byte("oi"))
	})

	httpServer.RegisterRoute(http.MethodGet, "/world", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("oi"))
	})

	if err := httpServer.Run(); err != nil {
		logger.Error("server error", logging.ErrorField(err))
	}
}
