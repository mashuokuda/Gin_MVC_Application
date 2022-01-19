package discuss

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/decree"
	"encoding/json"
	"gorm.io/gorm/clause"
	"hash"
)

type Discuss struct {
	Id           int           `gorm:"primaryKey;autoIncrement"`
	Decree       decree.Decree `gorm:"primaryKey;foreignKey:Id"`
	Create_User  int
	Discuss_Type int
	Title        string
	Opened       int
	Content      json.RawMessage `json:"content"`
}

type ContentJSON []struct {
	Title       string    `json:"title"`
	Hash        hash.Hash `json:"hash"`
	Create_User int       `json:"createUser"`
	Body        string    `json:"body"`
	MentionTo   []int     `json:"mentionTo"`
}

func GetDiscuss(id int) (Discuss, error) {
	var d Discuss
	err := database.DB.Find(&d, id, "id = ?").Error
	e := database.DB.Preload(clause.Associations).Find(&decree.Decree{}).Error
	if e != nil {
		err = e
	}
	return d, err
}

// func CreateDiscuss(discuss Discuss) error{
// 	var s string
// 	_ = json.NewEncoder().Encode(Discuss{})

// 	return database.DB.Create(&Discuss{
// 		Content: s,
// 	}).Error
// }
