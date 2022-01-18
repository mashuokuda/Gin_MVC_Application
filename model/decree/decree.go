package decree

import (
	"Gin_MVC/model/database"
	"os"
	"time"
)

type Decree struct {
	Id               int `gorm:"primaryKey;autoIncrement"`
	Decree_Reference string
	Name             string
	Last_update      *time.Time
}

func GetDecree(id int) (Decree, error) {
	decree := Decree{}
	err := database.DB.Find(&decree, "Id", id).Error
	return decree, err
}

func CreateDecree(decree Decree) error {
	return database.DB.Create(decree).Error
}

func (decree Decree) getDecree() (*os.File, error) {
	return os.Open(decree.Decree_Reference)
}
