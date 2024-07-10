package config

import (
	"fmt"
	"os"
	

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	SECRET_KEY string
	URL_PORT   string
}

func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}
	config.SECRET_KEY = cast.ToString(Coalesce("SECRET_KEY", "secret-key"))
	config.URL_PORT = cast.ToString(Coalesce("URL_PORT", "8081"))

	return config
}

func Coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
