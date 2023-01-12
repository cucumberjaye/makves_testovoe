package configs

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

var (
	Domain string
	Port   string
)

const DefaultPort = "8000"

func LoadConfig() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}
	Domain = mustEnvStr("DOMAIN")
	Port = optionalEnvStr("PORT", DefaultPort)

	return nil
}

func optionalEnvStr(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func mustEnvStr(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	} else {
		log.Panicf("Environment variable %v must be set.", key)
		return ""
	}
}
