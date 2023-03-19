package entity

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
)

func CreateNewMancalaGame() (*model.MancalaGame, error) {
	player1, err := CreateNewPlayer()
	if err != nil {
		return nil, err
	}

	player2, err := CreateNewPlayer()
	if err != nil {
		return nil, err
	}

	game := model.NewMancalaGame(player1, player2)

	result := config.DB.Create(&game)
	if result.Error != nil {
		return nil, result.Error
	}

	player1.MancalaGameID = game.ID
	player2.MancalaGameID = game.ID

	config.DB.Save(&player1)
	config.DB.Save(&player2)

	game.Player1 = *player1
	game.Player2 = *player2

	return game, nil
}

func GetMancalaGameById(game *model.MancalaGame, id string) error {
	result := config.DB.First(&game, id)

	player1 := config.DB.First(&game.Player1, game.Player1ID)
	if player1.Error != nil {
		return player1.Error
	}

	player2 := config.DB.First(&game.Player2, game.Player2ID)
	if player2.Error != nil {
		return player2.Error
	}

	if result.Error != nil {
		return result.Error
	}

	return nil
}
