package migrations

import (
	"dbo-test/database"
	"dbo-test/models"
)

func MigrateRangga() {

	db := database.DBConn

  DoMigration("20220918_MigrasiAwal", func() {
		db.AutoMigrate(&models.User{})
	})

	DoMigration("20220919_Migrasi", func() {
		db.AutoMigrate(&models.Catalog{})
	})

}
