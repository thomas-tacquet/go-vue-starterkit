package helpers

import (
	"fmt"

	"github.com/spf13/viper"
)

// DatabaseConfig is a struct to manage database configuration.
type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
	Schema   string
}

// InitWithViper simply takes a configured viper pointer and returns a new DatabaseConfig
// with hydrated fields according to viper's configuration.
func InitWithViper(vpr *viper.Viper) DatabaseConfig {
	return DatabaseConfig{
		Host:     vpr.GetString(EnvDBHost),
		Port:     vpr.GetString(EnvDBPort),
		User:     vpr.GetString(EnvDBUser),
		Password: vpr.GetString(EnvDBPassword),
		DBName:   vpr.GetString(EnvDBName),
		SSLMode:  vpr.GetString(EnvDBSSlMode),
		Schema:   vpr.GetString(EnvDBSchema),
	}
}

// ConnString returns a connexion string according to current configuration.
func (dbc DatabaseConfig) ConnString() string {
	const connexionString = "host=%s port=%s user=%s password=%s dbname=%s sslmode=%s search_path=%s"
	return fmt.Sprintf(connexionString, dbc.Host, dbc.Port, dbc.User, dbc.Password, dbc.DBName, dbc.SSLMode, dbc.Schema)
}
