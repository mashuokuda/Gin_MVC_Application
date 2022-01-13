package controller

import (
	"Gin_MVC/model/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	_ "net/http"
)

func DisplayLoginFrom(c *gin.Context) {
	session := sessions.Default(c)
	e := session.Get("err")
	if e != nil {
		log.Println("login failed")
		session.Delete("err")
	}
	c.HTML(200, "login.html", gin.H{
		"err": e,
	})
}

func DoAuth(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("err")
	if session.Get("user") != nil {
		c.Next()
	}
	postedUser := c.PostForm("username")
	user, _ := user.GetUser(postedUser)
	formPass := c.PostForm("password")
	log.Println(formPass + " : " + user.Password)
	p, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err := bcrypt.CompareHashAndPassword(p, []byte(formPass)); err != nil {
		log.Println("Cannot login")
		c.Abort()
		session.Set("err", "ログインに失敗しました")
		session.Save()
		c.Redirect(302, "/login")

	} else {
		session.Set("User", user.Username)
		session.Save()
		c.Next()
		//session.Set("length")
		log.Println("Logined User: " + user.Username)
		c.Redirect(302, "/")
	}
}
