package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const ()

// Config ...
type Config struct {
	Environment   string // develop, staging, production
	UserServiceHost string
	UserServicePort string
	ManageServiceHost string
	ManageServicePort string

	LogLevel string
	HTTPPort string
}

// Load loads environment vars and inflates Config
func Load() Config {

	if err := godotenv.Load(".env"); err != nil {
		fmt.Println("No .env file found")
	}
	config := Config{}

	config.Environment = cast.ToString(getOrReturnDefault("ENVIRONMENT", "prod"))
	config.LogLevel = cast.ToString(getOrReturnDefault("LOG_LEVEL", "debug"))
	config.HTTPPort = cast.ToString(getOrReturnDefault("HTTP_PORT", ":8080"))
	config.UserServiceHost = cast.ToString(getOrReturnDefault("USER_SERVICE_HOST", "localhost"))
	config.UserServicePort = cast.ToString(getOrReturnDefault("USER_GRPC_PORT", "7072"))
	config.ManageServiceHost = cast.ToString(getOrReturnDefault("MANAGE_SERVICE_HOST", "localhost"))
	config.ManageServicePort = cast.ToString(getOrReturnDefault("MANAGE_GRPC_PORT", "7073"))

	return config
}

func getOrReturnDefault(key string, defaultValue interface{}) interface{} {
	if os.Getenv(key) == "" {
		return defaultValue
	}

	return os.Getenv(key)
}
