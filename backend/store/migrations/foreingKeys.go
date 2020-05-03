package migrations

import (
	"reflect"

	"github.com/jinzhu/gorm"
)

//foreignKeyExecutor execute forein key actions
type foreignKeyExecutor func(*gorm.DB, interface{}, interface{}, *gorm.Relationship)

// AddForeignKeys must be called on tables to define forein key settings
func AddForeignKeys(db *gorm.DB, values ...interface{}) {
	autoForeignKey(db, executeAddForeignKey, values...)
}

// DropForeignKeys must be called on rollbacks when AddForeignKeys is called
func DropForeignKeys(db *gorm.DB, values ...interface{}) {
	autoForeignKey(db, executeDropForeignKey, values...)
}

// executeAddForeignKey adds CASCADE fk on table model
func executeAddForeignKey(db *gorm.DB, localTableModel interface{}, destTableModel interface{}, relationship *gorm.Relationship) {
	db.AutoMigrate(localTableModel, destTableModel)

	dest := db.NewScope(destTableModel).TableName() + "(" + relationship.AssociationForeignDBNames[0] + ")"
	db.Model(localTableModel).AddForeignKey(relationship.ForeignDBNames[0], dest, "CASCADE", "CASCADE")
}

// executeDropForeignKey drops a forein key
func executeDropForeignKey(db *gorm.DB, localTableModel interface{}, destTableModel interface{}, relationship *gorm.Relationship) {
	//	db.AutoMigrate(localTableModel, destTableModel)

	dest := db.NewScope(destTableModel).TableName() + "(" + relationship.AssociationForeignDBNames[0] + ")"
	db.Model(localTableModel).RemoveForeignKey(relationship.ForeignDBNames[0], dest)
}

//autoForeignKey automatisation de la détection de foreign keys dans les modèles de données donnés en paramètre
func autoForeignKey(db *gorm.DB, execute foreignKeyExecutor, values ...interface{}) {
	// for each struct
	for _, value := range values {
		scope := db.NewScope(value)

		// for each field
		for _, field := range scope.GetModelStruct().StructFields {
			if relationship := field.Relationship; relationship != nil {
				if _, ok := field.TagSettingsGet("polymorphic"); ok {
					// do not create fk on polymorphic keys
					continue
				}

				if relationship.Kind == "belongs_to" {

					if value, _ := field.TagSettingsGet("ASSOCIATION_AUTOCREATE"); value == "false" {
						// if fk creation is forced to false, do nothing
					} else {
						newValue := reflect.New(field.Struct.Type).Interface()
						execute(db, value, newValue, relationship)
					}
				}

				if relationship.Kind == "has_one" {
					newValue := reflect.New(field.Struct.Type).Interface()
					execute(db, newValue, value, relationship)
				}

				if relationship.Kind == "has_many" {
					newValue := reflect.New(field.Struct.Type.Elem()).Interface()
					execute(db, newValue, value, relationship)
				}
			}
		}
	}
}
