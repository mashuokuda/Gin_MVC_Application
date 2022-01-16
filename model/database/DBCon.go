package database

import (
	"errors"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() error {
	if DB == nil {
		var err error
		DB, err = gorm.Open(sqlite.Open("data.DB"), &gorm.Config{})

		if err != nil {
			return errors.New("Cannot connect database")
		}
	}
	return nil
}

func migrator(i interface{}) error {
	return DB.AutoMigrate(i)
}

func Migrator(i []interface{}) error {
	for _, i3 := range i {
		err := migrator(i3)
		if err != nil {
			return err
		}
	}
	return nil
}

func Transaction(f func(tx *gorm.DB) error) error {
	d := DB
	DB = DB.Begin()
	DB.Begin()
	err := DB.Transaction(f)
	if err != nil {
		DB.Rollback()
		log.Fatalln("Rollback")
		return err
	}
	DB = d
	return nil
}
