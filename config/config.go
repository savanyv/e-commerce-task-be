package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type database struct {
	HostDB     string
	PortDB     string
	UserDB     string
	PassDB     string
	NameDB     string
}

type Config struct {
	PG database
	PortServer string
}

func LoadConfig() *Config {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	return &Config{
		PG: database{
			HostDB:     getEnv("HOST_DB"),
			PortDB:     getEnv("PORT_DB"),
			UserDB:     getEnv("USER_DB"),
			PassDB:     getEnv("PASS_DB"),
			NameDB:     getEnv("NAME_DB"),
		},
		PortServer: getEnv("PORT_SERVER"),
	}
}

func getEnv(key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		return ""
	}

	return value
}
