package helpers

// ConfigPrefix is used to determine configuration's prefix, it's useful if you
// want a .env file with different configuration scopes.
const ConfigPrefix = "GOVUE"

// Now we define configuration's keys, you can use strings directly in your code but
// having them definied as const makes code clearer
const (
	EnvPort = "PORT"

	EnvPrivateRSA = "RSA_PRIVATE"
	EnvPublicRSA  = "RSA_PUBLIC"

	EnvLogPath  = "LOG_PATH"
	EnvLogName  = "LOG_NAME"
	EnvLogLevel = "LOG_LEVEL"

	EnvDBHost     = "DB_HOST"
	EnvDBPort     = "DB_PORT"
	EnvDBUser     = "DB_USER"
	EnvDBPassword = "DB_PASSWORD"
	EnvDBSSlMode  = "DB_SSLMODE"
	EnvDBName     = "DB_NAME"
	EnvDBSchema   = "DB_SCHEMA"
)

// System env variable
const (
	EnvApp = "GOVUE_ENV"
)
