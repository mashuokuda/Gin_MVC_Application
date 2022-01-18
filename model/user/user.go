package user

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/notify"
	"Gin_MVC/model/priority"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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
	Notify   notify.Notify     `gorm:"foreignKey:Id"`
	Priority priority.Priority `gorm:"foreignKey:Id"`
}

func GetUser(name string) (User, error) {
	//_ := database.DBConnection()
	user := User{}
	err := database.DB.Find(&user, "Username", name).Error
	if err == nil {
		database.DB.Preload(clause.Associations).Find(&user.Priority)
		database.DB.Preload(clause.Associations).Find(&user.Notify)
	}
	return user, err
}

func CreateUser(user *User, tx *gorm.DB) *gorm.DB {
	//_ := database.DBConnection()
	p, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(p)

	return tx.Create(user)
}
