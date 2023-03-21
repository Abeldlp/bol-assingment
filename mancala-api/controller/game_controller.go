package controller

import (
	"net/http"

	"github.com/Abeldlp/bol-assignment/mancala-api/entity"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
	"github.com/Abeldlp/bol-assignment/mancala-api/util"
	"github.com/gin-gonic/gin"
)

func GetAll(c *gin.Context) {
	games, err := entity.GetAllMancalaGames()

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": games})
}

func GetById(c *gin.Context) {
	var game model.MancalaGame
	if err := entity.GetMancalaGameById(&game, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	game.Player1.ID = game.Player1ID
	game.Player2.ID = game.Player2ID

	c.JSON(http.StatusOK, gin.H{"data": game})
}

func Create(c *gin.Context) {
	game, err := entity.CreateNewMancalaGame()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

func Update(c *gin.Context) {
	var game model.MancalaGame
	if err := entity.GetMancalaGameById(&game, c.Param("id")); err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}

	game.Player1.ID = game.Player1ID
	game.Player2.ID = game.Player2ID

	type Response struct {
		SelectedPit int `json:"selected-pit"`
	}

	var ResponseBody Response

	c.BindJSON(&ResponseBody)

	game.PlayRound(ResponseBody.SelectedPit)

	if err := entity.UpdateMancala(&game); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	util.LogBoard(game) //DEBUG PURPOSES

	c.JSON(http.StatusOK, game)
}
