package user

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/decree"
)

type Star struct {
	Id     int `gorm:"primaryKey"`
	UserId int
	Star   int
}

func (receiver Star) getStars() *[]decree.Decree {
	var d []decree.Decree
	database.DB.Find(&d, receiver.Star, "id = ?")
	return &d
}
