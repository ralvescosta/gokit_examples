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

	topology := rabbitmq.
		NewTopology().
		Exchange(rabbitmq.NewFanout("observability")).
		Queue(rabbitmq.NewQueue(QueueName).WithDql().WithRetry(3, 30*time.Second).Binding("observability", ""))

	client, err := rabbitmq.NewClient(params.Cfg, params.Logger).InstallTopology(topology)
	if err != nil {
		params.Logger.Error("could not start rabbitmq client", zap.Error(err))
	}

	dispatcher := rabbitmq.NewDispatcher(params.Logger, client, topology)

	params.BasicConsumer.Install(dispatcher)

	dispatcher.ConsumeBlocking(params.Sig)

	return nil
}

var ConsumerCmd = &cobra.Command{
	Use:   "consumer",
	Short: "Consumer Command",
	RunE:  RunCommand("consumer", consumer),
}
