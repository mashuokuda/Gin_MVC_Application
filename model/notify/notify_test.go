package notify

import (
	"Gin_MVC/model/database"
	"encoding/json"
	"testing"
)

func TestNotify(t *testing.T) {
	database.DBConnection()
	database.Migrator([]interface{}{&Notify{}})
	var j = NotifyJSON{
		struct {
			DiscussID int    "json:\"discussID\""
			Hash      string "json:\"hash\""
			Level     int    "json:\"level\""
			Comment   string "json:\"comment\""
		}{
			DiscussID: 0,
			Hash:      "",
			Level:     0,
			Comment:   "aa",
		},
	}
	s, _ := json.Marshal(j)
	var n = Notify{
		// UserID: 1,
		Notify: string(s),
	}
	err := CreateNotify(n.Id, database.DB)
	database.DB.Last(&n)
	// log.Panicln(n)
	for i := 0; i < 10; i++ {
		n, _ = n.AddNotify(
			NotifyJSON{
				struct {
					DiscussID int    "json:\"discussID\""
					Hash      string "json:\"hash\""
					Level     int    "json:\"level\""
					Comment   string "json:\"comment\""
				}{},
			},
		)
	}
	// var r = Notify{}
	// err = database.DB.First(&r, n.UserID).Update("Notify", n.Notify).Error

	if err.Error != nil {
		panic(err.Error)
	}

}
