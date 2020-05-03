package helpers

import (
	"fmt"

	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Schema   string
}

func InitWithViper(vpr *viper.Viper) DatabaseConfig {
	return DatabaseConfig{
		Host:     vpr.GetString("DB_HOST"),
		Port:     vpr.GetString("DB_PORT"),
		User:     vpr.GetString("DB_USER"),
		Password: vpr.GetString("DB_PASSWORD"),
		DBName:   vpr.GetString("DB_NAME"),
		SSLMode:  vpr.GetString("DB_SSLMODE"),
		Schema:   vpr.GetString("DB_SCHEMA"),
	}
}

func (dbc DatabaseConfig) String() string {
	const toBeFormatted = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s"
	return fmt.Sprintf(toBeFormatted, dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.DBName, dbc.SSLMode, dbc.Schema)
}
