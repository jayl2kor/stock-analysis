package database

import "gorm.io/gorm"

var DB *gorm.DB

func GetConnection() *gorm.DB {
	return DB
}
