package user

import (
	"log"
	"os"
)

//Image 画像のUUID.png
type Image string

/*
	GetImage
	画像パスを取得
	パスがなければデフォルト画像を表示
*/
func (receiver Image) GetImage() string {
	if receiver == "" {
		return defaultImage()
	}
	open, err := os.Open("resource/userResource/" + string(receiver))
	defer open.Close()
	if err != nil {
		log.Printf("File not found : %s", receiver)
		return defaultImage()
	}
	return "resource/userResource/" + string(receiver)
}

/*
	defaultImage
	デフォルト画像パスを渡す
*/
func defaultImage() string {
	var s Image
	s = "image.png"
	return s.GetImage()
}
