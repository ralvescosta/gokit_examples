package main

import (
	"github.com/ralvescosta/gokit_example/http_server/cmd"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var root = &cobra.Command{
	Use:     "app",
	Short:   "HTTP API Example",
	Version: "0.0.1",
}

func main() {
	root.AddCommand(cmd.ApiCmd)

	root.Execute()
}
