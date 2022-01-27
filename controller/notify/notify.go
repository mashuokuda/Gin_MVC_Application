package notify

import (
	"Gin_MVC/controller/login"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

//GetNotify 通知取得
func GetNotify(c *gin.Context) {
	user, _, err := login.GetLoginUser(c)
	log.Println("notify", err)
	if err != nil {
		c.JSON(403, "")
	} else {
		notifies := user.Notify
		log.Println(user)
		c.JSON(http.StatusAccepted, notifies.Notify)
	}
}
