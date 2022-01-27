package priority

import (
	"Gin_MVC/model/database"

	"gorm.io/gorm"
)

type Priority struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Priority int `gorm:"default:100;not null"`
}

func CreatePriority(id int, tx *gorm.DB) *gorm.DB {
	//d,_ := database.DBConnection()
	return tx.Create(&Priority{Id: id, Priority: 100})
}

func GetPriority(id int) *Priority {
	var pr Priority
	//d,_:= database.DBConnection()
	database.DB.Find(&pr, "id", id)
	return &pr
}
