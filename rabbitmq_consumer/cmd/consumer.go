package cmd

import (
	"os"
	"time"

	"github.com/ralvescosta/gokit/rabbitmq"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/consumers"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

const (
	QueueName = "observability.queue"
)

type ConsumerParams struct {
	CommonParams

	Sig           chan os.Signal
	BasicConsumer consumers.BasicConsumer
}

type Message struct {
	Value string `json:"value"`
}

func consumer(params ConsumerParams) error {
	params.Logger.Debug("Stating RabbitMQ Consumer...")

	channel, err := rabbitmq.NewChannel(params.Cfg.RabbitMQConfigs, params.Logger)
	if err != nil {
		params.Logger.Error("could not start rabbitmq client", zap.Error(err))
	}

	topology, err := rabbitmq.
		NewTopology(params.Logger).
		Exchange(rabbitmq.NewFanoutExchange("observability")).
		Queue(rabbitmq.NewQueue(QueueName).WithDQL().WithRetry(30*time.Second, 3)).
		QueueBinding(rabbitmq.NewQueueBinding().Queue(QueueName).Exchange("observability")).
		Channel(channel).
		Apply()
	if err != nil {
		params.Logger.Error("topology error", zap.Error(err))
	}

	dispatcher := rabbitmq.NewDispatcher(params.Logger, channel, topology.GetQueuesDefinition())

	params.BasicConsumer.Install(dispatcher)

	dispatcher.ConsumeBlocking(params.Sig)

	return nil
}

var ConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Consumer Command",
	RunE:  RunCommand("consumer", consumer),
}
