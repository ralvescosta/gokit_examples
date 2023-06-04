package cmd

import (
	"os"

	httpHandlers "github.com/ralvescosta/gokit_example/http_server_with_otlp/pkg/http_handlers"

	"github.com/ralvescosta/gokit/httpw"
	"github.com/spf13/cobra"
)

type APIParams struct {
	CommonParams

	Sig      chan os.Signal
	Handlers httpHandlers.HTTPHandlers
}

func api(params APIParams) error {
	params.Logger.Debug("Stating HTTP API...")

	router := httpw.
		NewHTTPServerBuilder(params.Cfg.HTTPConfigs, params.Logger).
		Signal(params.Sig).
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
