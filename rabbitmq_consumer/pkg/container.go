package pkg

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/ralvescosta/gokit/env"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit/tracing"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/consumers"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/internal/repositories"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/internal/services"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	container := dig.New()

	cfg, err := env.
		New().
		Tracing().
		HTTPServer().
		Messaging().
		Build()

	if err != nil {
		return nil, err
	}

	container.Provide(ProvideConfigs(cfg))
	container.Provide(logging.NewDefaultLogger)
	container.Invoke(InvokeJaegerTracingExporter)
	container.Provide(ProvideSignal)
	container.Provide(repositories.NewBookRepository)
	container.Provide(services.NewBookService)
	container.Provide(consumers.NewBasicConsumer)

	return container, nil
}

func ProvideConfigs(cfg *env.Config) func() *env.Config {
	return func() *env.Config {
		return cfg
	}
}

func InvokeJaegerTracingExporter(cfg *env.Config, logger logging.Logger) {
	tracing.NewJaeger(cfg, logger).Build(context.Background())
}

func ProvideSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return sig
}
