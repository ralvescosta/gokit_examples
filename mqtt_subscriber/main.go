package main

import (
	"os"

	configsBuilder "github.com/ralvescosta/gokit/configs_builder"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit/mqtt"
)

func main() {
	cfgs, err := configsBuilder.NewConfigsBuilder().MQTT().Build()
	if err != nil {
		panic(err)
	}

	logger, err := logging.NewDefaultLogger(cfgs.AppConfigs)
	if err != nil {
		panic(err)
	}

	client := mqtt.NewMQTTClient(cfgs, logger)
	if err := client.Connect(); err != nil {
		panic(err)
	}

	dispatcher := mqtt.NewDispatcher(logger, client.Client())

	_ = dispatcher.Register("ralvescosta/random/topic", mqtt.AtLeastOnce, func(topic string, qos mqtt.QoS, payload []byte) error {
		logger.Info("hello-world")
		return nil
	})

	signal := make(chan os.Signal)

	dispatcher.ConsumeBlocking(signal)
}
