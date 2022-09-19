package config

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	ISO_DATE           = "2006-01-02"
	ISO_DATE_TIME      = "2006-01-02 15:04:05"
	READABLE_DATE      = "02 Jan 2006"
	READABLE_DATE_TIME = "02 Jan 2006 15:04:05 -07"
	EMPTY_UUID         = "00000000-0000-0000-0000-000000000000"
)

func Env(key string, fallback string) string {

	godotenv.Load(".env")

	if value, exists := os.LookupEnv(key); exists {

		return value

	}

	envOS := os.Getenv(key)

	if envOS != "" {

		return envOS

	}

	if fallback != "" {

		return fallback

	}

	return ""

}
