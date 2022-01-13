package notify

import (
	"Gin_MVC/model/database"
	"encoding/json"
)

type Notify struct {
	UserID int    `gorm:"primaryKey"`
	Notify string `json:"notify"`
}

func CreateNotify(user int) error {
	var s string
	t, _ := json.Marshal(NotifyJSON{
		struct {
			DiscussID int    "json:\"discussID\""
			Hash      string "json:\"hash\""
			Level     int    "json:\"level\""
			Comment   string "json:\"comment\""
		}{},
	})
	s = string(t)
	return database.DB.Create(&Notify{
		UserID: user,
		Notify: s,
	}).Error
}

func GetNotify(user int) (Notify, error) {
	var n Notify
	e := database.DB.First(&n, user).Error
	return n, e
}

func (n Notify) AddNotify(comment NotifyJSON) (Notify, error) {
	var j NotifyJSON
	_ = json.Unmarshal([]byte(n.Notify), &j)
	j = append(j, comment...)
	b, _ := json.Marshal(j)
	n.Notify = string(b)
	var r = Notify{}
	tx := database.DB.First(&r, n.UserID).Update("Notify", n.Notify)
	database.DB.Find(&n, "user_id", r.UserID)
	return n, tx.Error
}

func (n Notify) RemoveNotify() error {
	var r = Notify{}
	b, _ := json.Marshal(NotifyJSON{})
	n.Notify = string(b)

	return database.DB.First(&r, n.UserID).Update("Notify", n.Notify).Error
}
