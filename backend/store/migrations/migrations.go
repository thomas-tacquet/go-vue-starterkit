package migrations

import (
	"gopkg.in/gormigrate.v1"
)

//Migrations List of migrations to do on startup
var Migrations = []*gormigrate.Migration{
	&V202005031209CreateUserTable,
}
