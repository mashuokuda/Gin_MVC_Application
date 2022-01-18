package user

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/decree"
)

type Star []int

func (receiver Star) getStars() *[]decree.Decree {
	var d []decree.Decree
	database.DB.Find(&d, receiver, "id = ?")
	return &d
}
