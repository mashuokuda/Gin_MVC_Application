package main

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/decree"
	"Gin_MVC/model/discuss"
	"Gin_MVC/model/notify"
	"Gin_MVC/model/priority"
	"Gin_MVC/model/user"
	"Gin_MVC/router"
)

//go:generate  air

func main() {
	_ = database.DBConnection()

	database.Migrator([]interface{}{&user.User{}, &decree.Decree{}, &notify.Notify{}, &priority.Priority{}, &discuss.Discuss{}})
	r := router.GetRouter()
	r.Run(":3000")
}
