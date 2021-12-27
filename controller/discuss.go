package controller

import "github.com/gin-gonic/gin"

func addDiscuss(c *gin.Context) {
	decree := c.Param("decree")
	if decree == "" {
		return
	}

}
