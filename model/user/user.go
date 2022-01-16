package user

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/notify"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       int `gorm:"primaryKey;autoIncrement"`
	Name     string
	Ruby     string
	Username string `gorm:"unique"`
	Password string
	Birth    *time.Time
	Tel      string
	Location uint32
	Publish  bool
	Profile  string
	Notify   notify.Notify
}

func GetUser(name string) (User, error) {
	//_ := database.DBConnection()
	user := User{}
	err := database.DB.Find(&user, "Username", name).Error
	return user, err
}

func CreateUser(user *User) error {
	//_ := database.DBConnection()
	p, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(p)
	return database.DB.Create(user).Error
}
