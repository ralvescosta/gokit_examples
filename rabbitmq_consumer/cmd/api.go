package cmd

import (
	"os"

	"github.com/ralvescosta/gokit_example/rabbitmq_consumer/pkg/handlers"
	"github.com/spf13/cobra"

	"github.com/ralvescosta/gokit/httpw/server"
)

type APIParams struct {
	CommonParams

	Sig      chan os.Signal
	Handlers handlers.HTTPHandlers
}

func api(params APIParams) error {
	params.Logger.Debug("Stating HTTP API...")

	router := server.
		NewHTTPServerBuilder().
		Configs(params.Cfg).
		Logger(params.Logger).
		Signal(params.Sig).
		WithTracing().
		WithMetrics().
		Build()

	params.Handlers.Install(router)

	return router.Run()
}

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "API Server Command",
	RunE:  RunCommand("api", api),
}
