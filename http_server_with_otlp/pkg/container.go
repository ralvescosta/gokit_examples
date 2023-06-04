package pkg

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ralvescosta/gokit/configs"
	configsBuilder "github.com/ralvescosta/gokit/configs_builder"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit/metrics"
	"github.com/ralvescosta/gokit/metrics/system"
	"github.com/ralvescosta/gokit/tracing"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/consumers"
	httpHandlers "github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/http_handlers"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/internal/repositories"
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/internal/services"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	container := dig.New()

	cfg, err := configsBuilder.NewConfigsBuilder().
		HTTP().
		Otel().
		Build()

	if err != nil {
		return nil, err
	}

	container.Provide(func() *configs.Configs { return cfg.(*configs.Configs) })
	container.Provide(func() *configs.AppConfigs { return cfg.(*configs.Configs).AppConfigs })
	container.Provide(logging.NewDefaultLogger)
	container.Invoke(InvokeOTLPTracingExporter)
	container.Invoke(InvokeMetricsExporter)
	container.Provide(ProvideSignal)
	container.Provide(repositories.NewBookRepository)
	container.Provide(services.NewBookService)
	container.Provide(httpHandlers.NewHandler)
	container.Provide(consumers.NewBasicConsumer)

	return container, nil
}

func InvokeOTLPTracingExporter(cfg *configs.Configs, logger logging.Logger) {
	tracing.NewOTLP(cfg, logger).
		WithApiKeyHeader().
		Build()
}

func InvokeMetricsExporter(cfg *configs.Configs, logger logging.Logger) {
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
