package user

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/discuss"
	"gorm.io/gorm/clause"

	"Gin_MVC/model/notify"
	"Gin_MVC/model/priority"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

/*
	User
	ユーザー情報を格納
*/
type User struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	//表示名
	Name string
	//フリガナ TODO:いる？
	Ruby string
	//ユーザーのID
	Username string `gorm:"unique"`
	//パスワードは作成時にハッシュ化
	Password string
	Birth    *time.Time
	Tel      string
	Location uint32
	Publish  bool
	Profile  string
	//画像はUUIDで一意に保存
	Image Image
	//Star　外部デーブル Has Many
	Star []Star
	//外部テーブル Has One
	Notify notify.Notify `gorm:"foreignKey:Id"`
	//外部テーブル Has One
	Priority priority.Priority `gorm:"foreignKey:Id"`
	//外部テーブル Has Many
	Discuss []discuss.Discuss `gorm:"foreignKey:Create_User"`
}

/*
	GetUser
	ユーザー情報をDBから取得
*/
func GetUser(name string) (User, error) {
	//_ := database.DBConnection()
	user := User{}
	err := database.DB.Find(&user, "Username", name).Error
	if err == nil {
		//DBから関連テーブルを取得
		database.DB.Preload(clause.Associations).Find(&user.Priority)
		database.DB.Preload(clause.Associations).Find(&user.Notify)
		database.DB.Preload(clause.Associations).Find(&user.Discuss)
		database.DB.Preload(clause.Associations).Find(&user.Star)

	}
	return user, err
}

/*
	CreateUser
	ユーザー作成
*/
func CreateUser(user *User, tx *gorm.DB) *gorm.DB {
	//_ := database.DBConnection()
	p, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	user.Password = string(p)

	return tx.Create(user)
}

/*
 */
func (user User) GetUserName() string {
	if user.Publish {
		return user.Name
	} else {
		return user.Username
	}
}
