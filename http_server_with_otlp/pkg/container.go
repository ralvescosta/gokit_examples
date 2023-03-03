package pkg

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ralvescosta/gokit/env"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit/metrics"
	"github.com/ralvescosta/gokit/metrics/system"
	"github.com/ralvescosta/gokit/tracing"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/consumers"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/handlers"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/internal/repositories"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/internal/services"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	container := dig.New()

	cfg, err := env.
		New().
		Otel().
		HTTPServer().
		Build()

	if err != nil {
		return nil, err
	}

	container.Provide(func() *env.Configs { return cfg })
	container.Provide(logging.NewDefaultLogger)
	container.Invoke(InvokeOTLPTracingExporter)
	container.Invoke(InvokeMetricsExporter)
	container.Provide(ProvideSignal)
	container.Provide(repositories.NewBookRepository)
	container.Provide(services.NewBookService)
	container.Provide(handlers.NewHandler)
	container.Provide(consumers.NewBasicConsumer)

	return container, nil
}

func InvokeOTLPTracingExporter(cfg *env.Configs, logger logging.Logger) {
	tracing.NewOTLP(cfg, logger).
		WithApiKeyHeader().
		Build()
}

func InvokeMetricsExporter(cfg *env.Configs, logger logging.Logger) {
	metrics.NewOTLP(cfg, logger).
		WithApiKeyHeader().
		Build()

	system.BasicMetricsCollector(logger)
}

func ProvideSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return sig
}
