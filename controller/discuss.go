package controller

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/discuss"
	"Gin_MVC/model/user"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
)

func addDiscuss(c *gin.Context) {
	decree := c.Param("decree")
	if decree == "" {
		return
	}
}

func Display(c *gin.Context) {
	discussId := c.Query("id")
	id, err := strconv.Atoi(discussId)
	if err != nil {
		log.Println("cannot parse ID :", discussId)
		c.Error(err)
	}
	dis, err := discuss.GetDiscuss(id)
	decname := dis.Decree.Name
	disTitle := dis.Title
	var u user.User
	er := database.DB.Find(&u, dis.Create_User, "id = ?").Error //Select * from user where user.id = Create_User
	if er != nil {
		u = user.User{
			Id:   -1,
			Name: "退会ユーザー",
		}
	}
	var content discuss.ContentJSON
	e := json.Unmarshal(dis.Content, &content)
	if e != nil {
	}
	c.HTML(200, "discuss.html", gin.H{
		"decree":     decname,
		"discuss":    disTitle,
		"createUser": u,
		"content":    content,
	})
}
