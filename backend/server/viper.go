package server

import (
	"os"

	"github.com/joho/godotenv"
)

// SetupViper is used to define config prefixes and set config file
// according to env variable
func (a *API) SetupViper() error {
	filename := ".env"

	switch os.Getenv("GOVUE_ENV") {
	case "production":
		filename = ".env.prod"
	}

	if err := godotenv.Overload(filename); err != nil {
		return err
	}

	a.Config.SetEnvPrefix("GOVUE")
	a.Config.AutomaticEnv()

	a.SetupViperDefaults()

	return nil
}

// SetupViperDefaults is used to define default necessary variables
func (a *API) SetupViperDefaults() {
}
