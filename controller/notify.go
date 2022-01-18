package controller

import (
	"Gin_MVC/model/user"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNotify(c *gin.Context) {
	user, err := user.GetUser("saitou")

	if err != nil {

	}
	//userprofile ,er :=
	//notifies, err := notify.GetNotify(user.Id)
	notifies := user.Notify
	log.Println(user)
	c.JSON(http.StatusAccepted, notifies.Notify)
}
