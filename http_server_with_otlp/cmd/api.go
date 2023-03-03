package cmd

import (
	"os"

	"github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/handlers"
	"github.com/spf13/cobra"

	"github.com/ralvescosta/gokit/httpw"
)

type APIParams struct {
	CommonParams

	Sig      chan os.Signal
	Handlers handlers.HTTPHandlers
}

func api(params APIParams) error {
	params.Logger.Debug("Stating HTTP API...")

	router := httpw.
		NewServer(params.Cfg.HTTPConfigs, params.Logger, params.Sig).
		WithTracing().
		Build()

	params.Handlers.Install(router)

	return router.Run()
}

var ApiCmd = &cobra.Command{
	Use:   "api",
	Short: "API Server Command",
	RunE:  RunCommand(api),
}
