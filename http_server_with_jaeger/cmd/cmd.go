package cmd

import (
	"github.com/ralvescosta/gokit/env"
	"github.com/ralvescosta/gokit/logging"
	"github.com/ralvescosta/gokit_example/http_server_with_jaeger/pkg"
	"github.com/spf13/cobra"
	"go.uber.org/dig"
	"go.uber.org/zap"
)

type CommonParams struct {
	dig.In

	Cfg    *env.Config
	Logger logging.Logger
}

func RunCommand(runner any) func(*cobra.Command, []string) error {
	return func(c *cobra.Command, s []string) error {
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
