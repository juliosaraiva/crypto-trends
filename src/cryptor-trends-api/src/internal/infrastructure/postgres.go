package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewConnection() {
	return gorm.Open(
		"postgres", "host=localhost port=5432 user=postgres dbname=cryptor_trends password=postgres sslmode=disable",
	)
}
