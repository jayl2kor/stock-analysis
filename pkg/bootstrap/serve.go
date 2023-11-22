package bootstrap

import (
	"stock-anaysis/cmd/stock_analysis/cli"
	"stock-anaysis/pkg/config"
	"stock-anaysis/pkg/database"
	"stock-anaysis/pkg/routing"
)

func Serve() {
	flags := cli.New()
	config.Set(flags.ConfigPath)

	database.Connect()

	routing.Init()
	routing.RegisterRoutes()
	routing.Serve()
}
