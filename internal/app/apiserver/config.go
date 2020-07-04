package apiserver

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

// Config ...
type Config struct {
	BindAddress string `toml:"bind_address"`
	LogLevel    string `toml:"log_level"`
	DatabaseURL string `toml:"database_url"`
	SessionKey  string `toml:"session_key"`
}

// viperEnvVariable loads db information from .env file
func viperEnvVariable(key string) string {
	viper.SetConfigFile("db.env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

// NewConfig ...
func NewConfig() *Config {
	// dbName := viperEnvVariable("POSTGRES_DB")
	// username := viperEnvVariable("POSTGRES_USER")
	// password := viperEnvVariable("POSTGRES_PASSWORD")
	dbName := "vaccinex_db"
	username := "hacker"
	password := "whosyourdaddy"
	pg_con_string := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", username, password, dbName)

	return &Config{
		BindAddress: ":8000",
		LogLevel:    "debug",
		DatabaseURL: pg_con_string,
	}
}
