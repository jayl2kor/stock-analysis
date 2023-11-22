package bootstrap

import (
	"stock-anaysis/cmd/stock_analysis/cli"
	"stock-anaysis/internal/database/migration"
	"stock-anaysis/pkg/config"
	"stock-anaysis/pkg/database"
)

func Migrate() {
	flags := cli.New()

	config.Set(flags.ConfigPath)
	database.Connect()

	// Auto migration
	migration.Migrate()
}
