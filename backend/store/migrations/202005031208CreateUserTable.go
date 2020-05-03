package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

type userV202005031209 struct {
	gorm.Model
	Nickname         *string                   `gorm:"unique_index"`
	Password         *string                   `gorm:"not null"`
	RememberMeTokens []rememberMeV202005031308 `gorm:"foreignkey:UserID"`
}

func (userV202005031209) TableName() string {
	return "users"
}

type rememberMeV202005031308 struct {
	gorm.Model
	User     *userV202005031209
	UserID   *uint   `gorm:"not null"`
	SeriesID *string `gorm:"not null;unique_index"`
	Token    *string `gorm:"not null"`
}

func (rememberMeV202005031308) TableName() string {
	return "remember_mes"
}

var V202005031209CreateUserTable = gormigrate.Migration{
	ID: "202005031209",
	Migrate: CreateTableMigration(
		&userV202005031209{},
		&rememberMeV202005031308{},
	),
	Rollback: func(tx *gorm.DB) error {
		DropForeignKeys(tx,
			&userV202005031209{},
			&rememberMeV202005031308{},
		)

		if tx.Error != nil {
			return tx.Error
		}

		return tx.DropTableIfExists(
			&userV202005031209{},
			&rememberMeV202005031308{},
		).Error
	},
}
