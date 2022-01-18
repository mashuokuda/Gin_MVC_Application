package router

import (
	"Gin_MVC/controller"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("graduate", store))
	//router.LoadHTMLGlob("view/*.html")
	router.LoadHTMLGlob("view/*/*.html")
	router.Static("resource", "./resource")

	router.GET("/", controller.IndexDisplayAction)
	router.GET("/discuss", controller.Display)
	router.GET("/login", controller.DisplayLoginFrom)
	router.POST("/doAuth", controller.DoAuth)
	router.GET("/notify", controller.GetNotify)
	router.GET("/profile", controller.ProfileDisplay)
	return router
}
