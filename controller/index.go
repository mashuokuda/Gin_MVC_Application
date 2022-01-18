package controller

import (
	"github.com/gin-gonic/gin"
)

func IndexDisplayAction(c *gin.Context) {
	errorMsg := ""
	usr, loginState, err := getLoginedUser(c)
	if err != nil {
		errorMsg = err.Error()
	}
	//userprofile ,er :=

	c.HTML(200, "index.html", gin.H{

		"username":    usr.Name,
		"userprofile": usr.Profile,
		"str":         "Index Page",
		"loginState":  loginState,
		"errorMsg":    errorMsg,
	})
}
