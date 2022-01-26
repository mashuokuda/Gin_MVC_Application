/*
	index.go
	indexページ処理
*/
package index

import (
	"Gin_MVC/controller/login"
	"github.com/gin-gonic/gin"
)

func IndexDisplayAction(c *gin.Context) {
	errorMsg := ""

	usr, loginState, err := login.GetLoginUser(c)
	if err != nil {
		errorMsg = err.Error()
	}

	c.HTML(200, "index.html", gin.H{
		//ユーザ名
		"username": usr.GetUserName(),
		//自己紹介
		"userprofile": usr.Profile,
		//アイコン
		"img": usr.Image.GetImage(),
		//ログイン状態
		"loginState": loginState,
		"errorMsg":   errorMsg,
	})
}
