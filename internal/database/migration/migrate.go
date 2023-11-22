package migration

import (
	"log"

	"stock-anaysis/pkg/database"
)

func Migrate() {
	db := database.GetConnection()

	err := db.AutoMigrate()

	if err != nil {
		panic(err)
	}

	log.Print("Migration Successfully Completed")

}
