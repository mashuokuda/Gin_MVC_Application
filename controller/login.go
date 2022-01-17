package controller

import (
	"Gin_MVC/model/user"
	"log"
	_ "net/http"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func DisplayLoginFrom(c *gin.Context) {
	session := sessions.Default(c)
	e := session.Get("err")
	if e != nil {
		log.Println("login failed")
		session.Delete("err")
		session.Save()
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
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(formPass)); err != nil {
		log.Println("Cannot login")
		c.Abort()
		session.Set("err", "ログインに失敗しました")
		session.Save()
		c.Redirect(302, "/login")

	} else {
		session.Set("User", user.Username)
		session.Set("LoginTime", time.Now().Unix())
		session.Save()
		c.Next()
		//session.Set("length")
		log.Println("Logined User: " + user.Username)
		c.Redirect(302, "/")
	}
}
