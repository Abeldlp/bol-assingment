package migration

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
)

func PlayerMigration() {
	db := config.GetDB()
	db.AutoMigrate(&model.Player{})
}
