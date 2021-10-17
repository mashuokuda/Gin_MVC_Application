package controller

import (
	"github.com/gin-gonic/gin"
)

func IndexDisplayAction(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"str": "Index Page",
	})
}
