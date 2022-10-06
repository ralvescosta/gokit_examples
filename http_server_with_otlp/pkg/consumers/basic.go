package consumers

import (
	"context"

	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit/rabbitmq"
	"go.uber.org/zap"
)

type (
	BasicConsumer interface {
		Install(dispatcher rabbitmq.Dispatcher)
	}

	basicConsumer struct {
		logger logging.Logger
	}
)

const (
	BasicQueueName = "observability.queue"
)

func (b *basicConsumer) Install(dispatcher rabbitmq.Dispatcher) {
	b.logger.Debug("Installing BasicConsumer...")
	dispatcher.RegisterDispatcher(BasicQueueName, BasicMessage{}, b.basicConsumer)
}

func (b *basicConsumer) basicConsumer(ctx context.Context, msg any, metadata *rabbitmq.DeliveryMetadata) error {
	basic := msg.(BasicMessage)

	b.logger.Info("Basic Consumer", zap.String("msg", basic.String()))
	return nil
}

func NewBasicConsumer(logger logging.Logger) BasicConsumer {
	return &basicConsumer{logger}
}
