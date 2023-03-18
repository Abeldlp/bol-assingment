package controller

import (
	"net/http"

	"github.com/Abeldlp/bol-assignment/mancala-api/entity"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
	"github.com/gin-gonic/gin"
)

func CreateMancalaGame(c *gin.Context) {
	game, err := entity.CreateNewMancalaGame()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": game})
}

func UpdateMancalaGame(c *gin.Context) {
	var game model.MancalaGame
	if err := entity.GetMancalaGameById(&game, c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// if err := c.BindJSON(&game); err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"message": "Could not serialize payload",
	// 	})
	// 	return
	// }

	// type Response struct {
	// 	SelectedPit int `json:"selectedPit"`
	// }

	// var ResponseBody Response

	// c.BindJSON(&ResponseBody)

	// game.PlayRound(ResponseBody.SelectedPit)

	// config.DB.Save(&game)
	// config.DB.Save(&game.Player1)
	// config.DB.Save(&game.Player2)

	c.JSON(http.StatusOK, game)
}
