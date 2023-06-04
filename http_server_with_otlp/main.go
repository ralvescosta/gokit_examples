package main

import (
	"github.com/ralvescosta/gokit_example/http_server_with_otlp/cmd"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var root = &cobra.Command{
	Use:     "app",
	Short:   "HTTP API Example",
	Version: "0.0.1",
}

// @title Go HTTP API
// @version 1.0
// @description This is a sample server.
// @termsOfService https://github.com/ralvescosta/gokit_examples/blob/main/LICENSE

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name MIT
// @license.url https://github.com/ralvescosta/gokit_examples/blob/main/LICENSE

// @host http://localhost:3333
// @BasePath /v1
func main() {
	root.AddCommand(cmd.ApiCmd)

	root.Execute()
}
