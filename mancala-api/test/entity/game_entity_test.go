package entity_test

import (
	"testing"

	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/entity"
	"github.com/Abeldlp/bol-assignment/mancala-api/migration"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
	"github.com/stretchr/testify/assert"
)

func TestGameEntity(t *testing.T) {
	config.SetupTestDB()
	migration.PlayerMigration()
	migration.GameMigration()
	defer config.CloseTestDB()

	t.Run("test get all mancala games", func(t *testing.T) {
		entity.CreateNewMancalaGame()
		entity.CreateNewMancalaGame()

		mancalaGames, err := entity.GetAllMancalaGames()

		assert.NoError(t, err)
		assert.NotNil(t, mancalaGames)
		assert.Equal(t, 2, len(mancalaGames))

	})

	t.Run("test create new game", func(t *testing.T) {
		entity.CreateNewMancalaGame()

		mancalaGames, err := entity.GetAllMancalaGames()

		assert.NoError(t, err)
		assert.NotNil(t, mancalaGames)
		assert.Equal(t, 3, len(mancalaGames))
	})

	t.Run("test get game by id", func(t *testing.T) {
		var game model.MancalaGame
		entity.GetMancalaGameById(&game, "1")

		assert.NotNil(t, game)
		assert.Equal(t, uint(0x1), game.ID)
	})

	t.Run("test update game", func(t *testing.T) {

		var game model.MancalaGame
		entity.GetMancalaGameById(&game, "1")

		game.Player1.Bucket = 3

		entity.UpdateMancala(&game)

		var updatedGame model.MancalaGame
		entity.GetMancalaGameById(&updatedGame, "1")

		assert.NotNil(t, updatedGame)
		assert.Equal(t, game.Player1.Bucket, updatedGame.Player1.Bucket)
	})
}
