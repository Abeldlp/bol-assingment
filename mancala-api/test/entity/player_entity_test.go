package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/entity"
	"github.com/Abeldlp/bol-assignment/mancala-api/migration"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
)

func TestPlayerEntity(t *testing.T) {
	config.SetupTestDB()
	migration.PlayerMigration()
	defer config.CloseTestDB()

	t.Run("test create new player", func(t *testing.T) {
		player, err := entity.CreateNewPlayer()
		assert.NoError(t, err)
		assert.NotNil(t, player)
		assert.NotZero(t, player.ID)
	})

	t.Run("test get player by id", func(t *testing.T) {
		newPlayer, _ := entity.CreateNewPlayer()

		player := &model.Player{}
		err := entity.GetPlayerById(player, newPlayer.ID)
		assert.NoError(t, err)
		assert.NotNil(t, player)
		assert.Equal(t, newPlayer.ID, player.ID)
	})

	t.Run("test update player", func(t *testing.T) {
		newPlayer, _ := entity.CreateNewPlayer()

		newPlayer.Bucket = 3

		err := entity.UpdatePlayer(newPlayer)
		assert.NoError(t, err)

		player := &model.Player{}
		err = entity.GetPlayerById(player, newPlayer.ID)
		assert.NoError(t, err)
		assert.NotNil(t, player)
		assert.Equal(t, newPlayer.ID, player.ID)
		assert.Equal(t, newPlayer.Bucket, player.Bucket)
	})
}
