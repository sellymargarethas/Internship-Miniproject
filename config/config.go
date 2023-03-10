package config

import (
	"Internship-Miniproject/constants"
	"os"

	"github.com/joho/godotenv"
)

var (
	DBDriver = GetEnv("DB_DRIVER", "postgres")
	DBName   = GetEnv("DB_NAME", "miniproject")
	DBHost   = GetEnv("DB_HOST", "localhost")
	DBPort   = GetEnv("DB_PORT", "5432")
	DBUser   = GetEnv("DB_USER", "postgres")
	DBPass   = GetEnv("DB_PASS", "admin")
	SSLMode  = GetEnv("SSL_MODE", "disable")

	APPUrl    = GetEnv("APP_URL")
	APPPort   = GetEnv("APP_PORT")
	APPPrefix = GetEnv("APP_PREFIX")

	JWT_KEY = GetEnv("JWT_KEY")
)

func GetEnv(key string, value ...string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error Load file .env not found")
	}

	if os.Getenv(key) != constants.EMPTY_VALUE {
		return os.Getenv(key)
	} else {
		if len(value) > constants.EMPTY_VALUE_INT {
			return value[constants.EMPTY_VALUE_INT]
		}
		return ""
	}
}
