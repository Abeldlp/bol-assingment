package controller_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/Abeldlp/bol-assignment/mancala-api/config"
	"github.com/Abeldlp/bol-assignment/mancala-api/controller"
	"github.com/Abeldlp/bol-assignment/mancala-api/entity"
	"github.com/Abeldlp/bol-assignment/mancala-api/migration"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGameController(t *testing.T) {
	config.SetupTestDB()
	migration.PlayerMigration()
	migration.GameMigration()
	defer config.CloseTestDB()

	t.Run("test get all mancala games", func(t *testing.T) {
		router := gin.Default()
		req, _ := http.NewRequest("GET", "/mancala-game", nil)
		res := httptest.NewRecorder()

		entity.CreateNewMancalaGame()

		router.GET("/mancala-game", controller.GetAll)
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		var response struct {
			Data []model.MancalaGame `json:"data"`
		}
		if err := json.Unmarshal(res.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}

		assert.Greater(t, len(response.Data), 0)
	})

	t.Run("test get mancala game by id", func(t *testing.T) {
		router := gin.Default()
		req, _ := http.NewRequest("GET", "/games/1", nil)
		res := httptest.NewRecorder()

		game, _ := entity.CreateNewMancalaGame()

		router.GET("/mancala-game/:id", controller.GetById)
		req.URL.Path = "/mancala-game/" + strconv.FormatUint(uint64(game.ID), 10)
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		var response struct {
			Data model.MancalaGame `json:"data"`
		}
		if err := json.Unmarshal(res.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, strconv.FormatUint(uint64(game.ID), 10), strconv.FormatUint(uint64(game.ID), 10))
	})

	t.Run("test create new mancala game", func(t *testing.T) {
		router := gin.Default()
		req, _ := http.NewRequest("POST", "/mancala-game", nil)
		res := httptest.NewRecorder()

		router.POST("/mancala-game", controller.Create)
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusCreated, res.Code)

		var response struct {
			Game model.MancalaGame `json:"data"`
		}

		if err := json.Unmarshal(res.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}

		assert.NotNil(t, response.Game)
		assert.Equal(t, 0, response.Game.Turn)
		assert.Equal(t, 0, response.Game.Player1.Bucket)
	})

	t.Run("test update mancala game", func(t *testing.T) {
		router := gin.Default()
		jsonData := `{"selected-pit": 3}`
		req, _ := http.NewRequest("PUT", "/mancala-game/1", strings.NewReader(jsonData))
		res := httptest.NewRecorder()

		game, _ := entity.CreateNewMancalaGame()

		req.URL.Path = "/mancala-game/" + strconv.FormatUint(uint64(game.ID), 10)

		router.PUT("/mancala-game/:id", controller.Update)
		router.ServeHTTP(res, req)

		assert.Equal(t, http.StatusOK, res.Code)

		var response struct {
			ID      uint         `json:"ID"`
			Turn    int          `json:"turn"`
			Player1 model.Player `json:"Player1"`
			Player2 model.Player `json:"Player2"`
		}
		if err := json.Unmarshal(res.Body.Bytes(), &response); err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, game.ID, response.ID)
		assert.Equal(t, 1, response.Turn)
		assert.Equal(t, 1, response.Player1.Bucket)
		assert.Equal(t, 5, response.Player2.Holes[5])
	})
}
