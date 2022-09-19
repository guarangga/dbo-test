package migrations

import (
	"fmt"
	"dbo-test/database"
	"dbo-test/models"
  "time"
)

func Migrate() {

	db := database.DBConn

	fmt.Println("Migrating...")

	db.AutoMigrate(&models.Migration{})

  MigrateRangga()

}

type fn func()

func DoMigration(key string, funcMigrate fn) {

	db := database.DBConn

	var migration models.Migration

	if count := db.Where(&models.Migration{Migration: key}).First(&migration).RowsAffected; count == 0 {

		funcMigrate()

		migration.Migration = key

		migration.Batch = int(time.Now().Unix())

		defer db.Create(&migration)

	}
}
