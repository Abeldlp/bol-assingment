package entity

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
	"gorm.io/gorm/clause"
)

func GetAllMancalaGames() ([]model.MancalaGame, error) {
	var games []model.MancalaGame

	result := config.DB.Preload(clause.Associations).Order("ID asc").Find(&games)
	if result.Error != nil {
		return nil, result.Error
	}

	return games, nil
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

	UpdatePlayer(player1)
	UpdatePlayer(player2)

	game.Player1 = *player1
	game.Player2 = *player2

	return game, nil
}

func UpdateMancala(game *model.MancalaGame) error {
	p1err := UpdatePlayer(&game.Player1)
	if p1err != nil {
		return p1err
	}

	p2err := UpdatePlayer(&game.Player2)
	if p2err != nil {
		return p2err
	}

	result := config.DB.Omit(clause.Associations).Save(&game)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
