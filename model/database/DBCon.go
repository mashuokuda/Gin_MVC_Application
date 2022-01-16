package database

import (
	"errors"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() error {
	if DB == nil {
		var err error
		DB, err = getNewDB()

		if err != nil {
			return errors.New("Cannot connect database")
		}
	}
	return nil
}

func getNewDB() (*gorm.DB, error) {
	var db *gorm.DB
	db, err := gorm.Open(sqlite.Open("data.DB"), &gorm.Config{})

	if err != nil {
		return nil, err
	}
	return db, nil
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
	//db, _ := getNewDB()
	return DB.Transaction(f)
}
