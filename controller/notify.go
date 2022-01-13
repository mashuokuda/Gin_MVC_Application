package controller

import (
	"Gin_MVC/model/notify"
	"Gin_MVC/model/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetNotify(c *gin.Context) {
	user, err := user.GetUser("saitou")

	if err != nil {

	}
	//userprofile ,er :=
	notifies, err := notify.GetNotify(user.Id)
	c.JSON(http.StatusAccepted, notifies.Notify)
}
