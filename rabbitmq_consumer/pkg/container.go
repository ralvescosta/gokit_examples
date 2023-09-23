package pkg

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ralvescosta/gokit/configs"
	configsBuilder "github.com/ralvescosta/gokit/configs_builder"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit/tracing"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/consumers"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/internal/repositories"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/internal/services"
	"go.uber.org/dig"
)

func NewContainer() (*dig.Container, error) {
	container := dig.New()

	cfg, err := configsBuilder.
		NewConfigsBuilder().
		Otel().
		HTTP().
		RabbitMQ().
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

func ProvideConfigs(cfg *configs.Configs) func() *configs.Configs {
	return func() *configs.Configs {
		return cfg
	}
}

func InvokeJaegerTracingExporter(cfg *configs.Configs, logger logging.Logger) {
	tracing.NewJaegerBuilder().Configs(cfg).Logger(logger).Build()
}

func ProvideSignal() chan os.Signal {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return sig
}
