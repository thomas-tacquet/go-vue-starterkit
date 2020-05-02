package store

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gopkg.in/gormigrate.v1"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/qor/validations"

	"github.com/thomas-tacquet/go-vue-starterkit/backend/helpers"
	migrations "github.com/thomas-tacquet/go-vue-starterkit/backend/store/migrations"
)

//InitAndGetDB init database connexion and returns it
func InitAndGetDB(reset bool, schema string, eLogrus *logrus.Entry) *gorm.DB {
	myOptions := gormigrate.DefaultOptions

	// use database transactions only if reset is not set to true
	if !reset {
		myOptions.UseTransaction = true
	}

	db := CreateDBInstance(schema, eLogrus)
	m := gormigrate.New(db, myOptions, migrations.Migrations)

	if reset {
		var err error
		for ; err == nil; err = m.RollbackLast() {
		}

		if err != gormigrate.ErrNoRunMigration {
			log.Printf("Error during rollback : %s", err)
		} else {
			log.Printf("Database has been cleaned")
		}

	}

	if err := m.Migrate(); err != nil {
		log.Fatalf("could not migrate : %v", err)
	}

	log.Printf("Migrations ran successfully")

	validations.RegisterCallbacks(db)

	log.Printf("Added validation callbacks")

	return db
}

func CreateDBInstance(schema string, eLogrus *logrus.Entry) *gorm.DB {

	dbConnexionString := helpers.DatabaseFormat(
		"127.0.0.1",
		"5431",
		"vuego",
		"vuego123+",
		"vuego",
		"disable",
		schema)

	var err error
	var db *gorm.DB

	if db, err = gorm.Open("postgres", dbConnexionString); err != nil {
		log.Printf("Error initializing db on 5433 : %v", err)
	}

	if err := db.DB().Ping(); err != nil {
		log.Fatalf("Error pinging db : %v", err)
	}

	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(25)

	if eLogrus != nil {
		db.SetLogger(eLogrus)
	} else {
		db.SetLogger(
			DefaultDatabaseLogger{
				log.New(os.Stdout, "\r\n", 0),
			},
		)
	}

	db.LogMode(true)

	log.Printf("Connected to DB successfully")
	return db
}

func GetDB(c *gin.Context) *gorm.DB {
	return c.MustGet("DB").(*gorm.DB)
}

func GetDBOverride(c *gin.Context) *gorm.DB {
	return c.MustGet("DBOverride").(*gorm.DB)
}

type DefaultDatabaseLogger struct {
	gorm.LogWriter
}

func (logger DefaultDatabaseLogger) Print(values ...interface{}) {
	messages := gorm.LogFormatter(values...)
	var out []interface{}
	for _, currM := range messages {
		switch v := currM.(type) {
		case string:
			if len(v) > 500 {
				out = append(out, fmt.Sprintf("%.500s...[truncated]", v))
			} else {
				out = append(out, v)
			}
		default:
			out = append(out, v)
		}
	}
	logger.Println(out)
}
