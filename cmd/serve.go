package cmd

import (
	"github.com/spf13/cobra"
	"stock-anaysis/pkg/bootstrap"
)

func init() {
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Serve app on dev server",
	Long:  `Application will be served on host and port defined in config.yaml file`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func serve() {
	bootstrap.Serve()
}
