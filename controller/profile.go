package controller

import (
	"Gin_MVC/model/decree"
	"github.com/gin-gonic/gin"
)

func ProfileDisplay(c *gin.Context) {
	var decName []decree.Decree
	errorMsg := ""
	usr, loginState, err := getLoginedUser(c)
	if err != nil {
		errorMsg = err.Error()
	}
	for _, star := range usr.Star {
		getDecree, err := decree.GetDecree(star.Star)
		if err != nil {
			return
		}
		decName = append(decName, getDecree)
	}
	c.HTML(200, "profile.html", gin.H{

		"username":    usr.Name,
		"userprofile": usr.Profile,
		"stars":       decName,
		"img":         usr.Image.GetImage(),
		"loginState":  loginState,
		"errorMsg":    errorMsg,
	})
}
