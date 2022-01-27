package user

import (
	"Gin_MVC/model/database"
	"Gin_MVC/model/decree"
	"gorm.io/gorm"
)

/*
	Star
	お気に入りリスト
*/
type Star struct {
	Id     int `gorm:"primaryKey"`
	UserId int
	Star   int
}

/*
	getStars
	お気に入りから法令リストを取得
*/
func (receiver Star) getStars() *[]decree.Decree {
	var d []decree.Decree
	database.DB.Find(&d, receiver.Star, "id = ?")
	return &d
}

/*
	AppendStar
	ユーザーにお気に入りを追加
*/
func (u *User) AppendStar(decree int) {
	s := Star{
		UserId: u.Id,
		Star:   decree,
	}
	database.Transaction(func(tx *gorm.DB) error {
		var err error
		if err = tx.Create(&s).Error; err != nil {
			return err
		}
		u.Star = append(u.Star, s)
		if err = tx.Save(&u).Error; err != nil {
			return err
		}
		return nil
	})
	database.DB.Save(&u)
}
