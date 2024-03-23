package dao

import (
	"github.com/e421083458/golang_common/lib"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() *gorm.DB {
	return db
}

func Migrate() error {
	return db.Migrator().AutoMigrate(
		&User{},
		&StoneTypes{},
		&Stone{},
	)
}

func InitTx() error {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		return err
	}
	db = tx
	return Migrate()
}
