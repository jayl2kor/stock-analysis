package config

import (
	"fmt"
)

type Config struct {
	App    App
	Server Server
	DB     DB
}

type App struct {
	Name string `yaml:"name"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type DB struct {
	Driver string `yaml:"driver"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Database string `yaml:"database"`

	SQLite3FileName string `yaml:"sqlite3filename"`
}

func (db *DB) DSN() string {
	switch db.Driver {
	case "sqlite3":
		return fmt.Sprintf("%s?_foreign_keys=on", db.SQLite3FileName)
	case "mysql":
		if db.Host == "" || db.Port == "" || db.Username == "" || db.Password == "" || db.Database == "" {
			return ""
		}
	}
	return ""
}
