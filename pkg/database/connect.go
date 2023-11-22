package database

import (
	"stock-anaysis/pkg/config"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	config := config.Get()

	dialector := newDialector(config.DB.Driver, config.DB.DSN())
	gormConfig := getDefaultGormConfig()

	gormConn, err := gorm.Open(dialector, gormConfig)
	if err != nil {
		panic("Cannot connect to database")
	}

	DB = gormConn
}

func newDialector(name, dsn string) gorm.Dialector {
	switch name {
	case "mysql":
		return mysql.Open(dsn)
	case "sqlite3":
		return sqlite.Open(dsn)
	}
	return nil
}

func getDefaultGormConfig() *gorm.Config {
	return &gorm.Config{}
}
