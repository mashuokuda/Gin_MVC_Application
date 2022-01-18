package controller

import (
	"Gin_MVC/model/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"log"
)

func IndexDisplayAction(c *gin.Context) {
	session := sessions.Default(c)
	sessionUser := session.Get("User") //will be nil
	usr := user.User{}
	loginState := false
	errorMsg := ""
	if sessionUser != nil {
		var err error
		usr, err = user.GetUser(sessionUser.(string))
		loginState = true
		if err != nil {
			log.Println("login error")
			usr = user.User{}
			loginState = false
			errorMsg = "不明なエラーが発生しました"
			session.Clear()
		}
	}
	//userprofile ,er :=

	c.HTML(200, "index.html", gin.H{

		"username":    usr.Name,
		"userprofile": usr.Profile,
		"str":         "Index Page",
		"loginState":  loginState,
		"errorMsg":    errorMsg,
	})
}
