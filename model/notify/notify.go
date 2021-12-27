package notify

import (
	"Gin_MVC/model/database"
	"encoding/json"
)

type Notify struct {
	UserID int    `gorm:"primaryKey"`
	Notify string `json:"notify"`
}

type NotifyJSON []struct {
	DiscussID int    `json:"discussID"`
	Hash      string `json:"hash"`
	Level     int    `json:"level"`
	Comment   string `json:"comment"`
}

func InitNotify(user int) error {
	var s string
	t, _ := json.Marshal(NotifyJSON{})
	s = string(t)
	return database.DB.Create(&Notify{
		UserID: user,
		Notify: s,
	}).Error
}
