package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotify(c *gin.Context) {
	user, _, err := getLoginedUser(c)
	log.Println("notify", err)
	if err != nil {
		c.JSON(403, "")
	} else {
		notifies := user.Notify
		log.Println(user)
		c.JSON(http.StatusAccepted, notifies.Notify)
	}
}
