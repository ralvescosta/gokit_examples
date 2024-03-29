package cmd

import (
	"os"

	"github.com/ralvescosta/gokit/configs"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit_example/api_and_consumer/pkg"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type CommonParams struct {
	dig.In

	Cfg    *configs.Configs
	Logger logging.Logger
}

func RunCommand(cmd string, runner any) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, args []string) error {
		os.Setenv("APP_NAME", cmd)

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
