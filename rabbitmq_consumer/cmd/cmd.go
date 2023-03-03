package cmd

import (
	"github.com/ralvescosta/gokit/env"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type CommonParams struct {
	dig.In

	Cfg    *env.Configs
	Logger logging.Logger
}

func RunCommand(cmd string, runner any) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, args []string) error {
		ioc, err := pkg.NewContainer()
		if err != nil {
			return err
		}

		err = ioc.Invoke(func(logger logging.Logger) error {
			if err := ioc.Invoke(runner); err != nil {
				logger.Error("error running command", zap.Error(err))
				return err
			}

			return nil
		})

		return err
	}
}
