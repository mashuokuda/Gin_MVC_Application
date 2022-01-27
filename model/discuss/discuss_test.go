package discuss

import (
	"Gin_MVC/model/database"
	"testing"
)

func TestDiscuss(t *testing.T) {
	_ = database.DBConnection()

	database.Migrator([]interface{}{&Discuss{}})
}
