package user

import (
	"log"
	"os"
)

type Image string

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

func defaultImage() string {
	var s Image
	s = "image.png"
	return s.GetImage()
}
