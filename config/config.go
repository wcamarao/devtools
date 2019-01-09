package config

import (
	"fmt"
	"os"
)

var c *Config

// Config contains service configuration
type Config struct {
	DB *DB
}

// DB contains database configuration
type DB struct {
	Host    string
	User    string
	Pass    string
	Name    string
	SSLMode string
}

// GetConfig fetches the service configuration from environment variables
func GetConfig() *Config {
	if c == nil {
		c = &Config{
			DB: getDB(),
		}
	}
	return c
}

func getDB() *DB {
	return &DB{
		Host:    get("DB_HOST"),
		User:    get("DB_USER"),
		Pass:    get("DB_PASS"),
		Name:    get("DB_NAME"),
		SSLMode: get("DB_SSLMODE"),
	}
}

func get(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("missing environment variable: %s", key))
	}
	return val
}
