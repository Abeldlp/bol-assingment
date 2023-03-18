package migrations

import (
	"github.com/Abeldlp/bol-assignment/user-api/model"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&model.MancalaUser{})
}
