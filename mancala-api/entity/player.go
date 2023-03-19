package entity

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
)

func CreateNewPlayer() (*model.Player, error) {
	player := model.NewPlayer()

	result := config.DB.Create(&player)
	if result.Error != nil {
		return nil, result.Error
	}

	return player, nil
}

func GetPlayerById(player *model.Player, id string) error {
	result := config.DB.First(&player, id)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func UpdatePlayer(player *model.Player, id int) error {
	result := config.DB.First(&player, id)
	if result.Error != nil {
		return result.Error
	}

	result = config.DB.Save(&player)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
