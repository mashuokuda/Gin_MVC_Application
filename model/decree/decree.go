/*
	decree
	法令
*/
package decree

import (
	"Gin_MVC/model/database"
	"os"
	"time"
)

type Decree struct {
	Id int `gorm:"primaryKey;autoIncrement"`
	//gitディレクトリのPath
	Decree_Reference string
	//法令名
	Name string
	//Last Commit Time
	Last_update *time.Time
}

/*
	法令取得
*/
func GetDecree(id int) (Decree, error) {
	decree := Decree{}
	err := database.DB.Find(&decree, "Id", id).Error
	return decree, err
}

/*
	法令作成

*/
func CreateDecree(decree Decree) error {
	//TODO: バッチ処理でこれを呼び出す
	return database.DB.Create(decree).Error
}

func (decree Decree) getDecree() (*os.File, error) {
	return os.Open(decree.Decree_Reference)
}
