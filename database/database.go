package database

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"dbo-test/config"
)

var (
	DBConn *gorm.DB
)

func InitDB() {
	var err error
	dsn := "host=" + config.Env("DB_HOST", "localhost") + " user=" + config.Env("DB_USERNAME", "postgres") + " password=" + config.Env("DB_PASSWORD", "1234") + " dbname=" + config.Env("DB_DATABASE", "dbo_test") + " port=" + config.Env("DB_PORT", "5432") + " sslmode=" + config.Env("DB_SSLMODE", "disable") + " TimeZone=" + config.Env("APP_TIMEZONE", "Asia/Jakarta")
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		 Logger: logger.Default.LogMode(logger.Silent),
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "",
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Established")
}
