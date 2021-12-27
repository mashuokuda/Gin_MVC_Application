package priority

import "Gin_MVC/model/database"

type Priority struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Priority int `gorm:"default:100;not null"`
}

func CreatePriority(id int) {
	//d,_ := database.DBConnection()
	database.DB.Create(&Priority{Id: id, Priority: 100})
}

func GetPriority(id int) *Priority {
	var pr Priority
	//d,_:= database.DBConnection()
	database.DB.Find(&pr, "id", id)
	return &pr
}
