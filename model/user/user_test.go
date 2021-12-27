package user

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/notify"
	"Gin_MVC/model/priority"
	"log"
	"testing"
)

func TestUser(t *testing.T) {
	_ = database.DBConnection()

	database.Migrator([]interface{}{&User{}, &notify.Notify{}, &priority.Priority{}})
	log.Println(GetUser("saitou").Name)

	var u = User{
		//UserId:       0,
		Name:     "斉藤",
		Ruby:     "サイトウ",
		Username: "saito",
		Password: "asdfgdf",
		Tel:      "03000000000",
		Location: 0,
		Publish:  false,
	}
	err := UserCreate(&u)
	if err != nil {
		return
	}
	err = notify.InitNotify(u.Id)
	if err != nil {
		return
	}
	priority.CreatePriority(u.Id)

}
