package pkg

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ralvescosta/gokit/env"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit/metric"
	"github.com/ralvescosta/gokit/metric/basic"
	"github.com/ralvescosta/gokit/tracing"
	"github.com/ralvescosta/gokit_example/http_server/pkg/consumers"
	"github.com/ralvescosta/gokit_example/http_server/pkg/handlers"
	"github.com/ralvescosta/gokit_example/http_server/pkg/internal/repositories"
	"github.com/ralvescosta/gokit_example/http_server/pkg/internal/services"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	container := dig.New()

	cfg, err := env.
		New().
		Tracing().
		HTTPServer().
		Build()

	if err != nil {
		return nil, err
	}

	container.Provide(func() *env.Config { return cfg })
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

func InvokeOTLPTracingExporter(cfg *env.Config, logger logging.Logger) {
	tracing.NewOTLP(cfg, logger).
		WithApiKeyHeader().
		Build()
}

func InvokeMetricsExporter(cfg *env.Config, logger logging.Logger) {
	metric.NewOTLP(cfg, logger).
		WithApiKeyHeader().
		Build()

	go basic.BasicMetricsCollector(logger, 15)
}

func ProvideSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return sig
}
