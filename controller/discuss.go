package controller

import (
	"Gin_MVC/model/notify"
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
	a, _ := user.GetUser("saitou")
	b, _ := notify.GetNotify(a.Id)
	var j notify.NotifyJSON
	json.Unmarshal([]byte(b.Notify), &j)
	log.Println(a.Name)
	c.HTML(200, "index.html", gin.H{
		"str": "ようこそ" + a.Name + "さん" + j[0].Comment,
	})
}
