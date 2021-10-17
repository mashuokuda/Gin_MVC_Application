package router

import (
	"Gin_MVC/controller"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	router.Static("resource", "./resource")

	router.GET("/", controller.IndexDisplayAction)

	return router
}
