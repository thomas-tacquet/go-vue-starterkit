package server

import (
	"os"

	"github.com/joho/godotenv"

	"github.com/thomas-tacquet/go-vue-starterkit/backend/helpers"
)

// SetupViper is used to define config prefixes and set config file
// according to env variable
func (a *API) SetupViper() error {
	filename := ".env"

	switch os.Getenv(helpers.EnvApp) {
	case "production":
		filename = ".env.prod"
	}

	if err := godotenv.Overload(filename); err != nil {
		return err
	}

	a.Config.SetEnvPrefix(helpers.ConfigPrefix)
	a.Config.AutomaticEnv()

	a.SetupViperDefaults()

	return nil
}

// SetupViperDefaults is used to define default necessary variables
func (a *API) SetupViperDefaults() {
}
