package migrations

import (
	"github.com/jinzhu/gorm"
	"github.com/thoas/go-funk"
)

type tableInterface interface {
	TableName() string
}

func CreateTableMigration(tables ...tableInterface) func(tx *gorm.DB) error {
	var itables []interface{}

	funk.ConvertSlice(tables, &itables)

	return func(tx *gorm.DB) error {
		if err := tx.AutoMigrate(itables...).Error; err != nil {
			return err
		}

		AddForeignKeys(tx,
			itables...,
		)
		return tx.Error
	}
}
