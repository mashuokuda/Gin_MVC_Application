package profile

import (
	"Gin_MVC/controller/login"
	_ "embed"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

//go:embed LocationList.json
var locJson []byte

func EditProfileDisplay(c *gin.Context) {
	errorMsg := ""
	usr, loginState, err := login.GetLoginUser(c)
	var locList []string
	json.Unmarshal(locJson, &locList)
	if err != nil {
		errorMsg = err.Error()
	}

	c.HTML(200, "editProfile.html", gin.H{
		"user": struct {
			//本名
			Name string
			//ユーザー名
			UserName string
			//自己紹介
			Profile string
			//居住地
			Location uint32
			//公開設定
			Publish bool
			//電話番号
			Tel string
		}{
			usr.Name,
			usr.Username,
			usr.Profile,
			usr.Location,
			usr.Publish,
			usr.Tel,
		}, //パスワード等を秘匿
		//アイコン
		"img": usr.Image.GetImage(),
		//ログイン状態
		"loginState":   loginState,
		"errorMsg":     errorMsg,
		"LocationList": locList,
	})
}
