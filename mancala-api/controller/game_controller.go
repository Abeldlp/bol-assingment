package controller

import (
	"net/http"

	"github.com/Abeldlp/bol-assignment/mancala-api/entity"
	"github.com/Abeldlp/bol-assignment/mancala-api/model"
	"github.com/Abeldlp/bol-assignment/mancala-api/util"
	"github.com/gin-gonic/gin"
)

// ListMancalaGames godoc
//
//	@Summary		List mancala games
//	@Description	get mancala games
//	@Produce		json
//	@Success		200	{array}		model.MancalaGame
//	@Failure		404	{object}	object{error string}
//	@Router			/mancala-game [get]
//	@Tags			MancalaGames
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

// ShowMancalaGame godoc
//
//	@Summary		Get mancala game by id
//	@Description	get single mancala game
//	@Produce		json
//	@Param			id	path		int	true	"MancalaGame ID"
//	@Success		200	{object}	model.MancalaGame
//	@Failure		404
//	@Router			/mancala-game/{id} [get]
//	@Tags			MancalaGames
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

// CreateMancalaGame godoc
//
//	@Summary		Create new Mancala along with two players
//	@Description	Create new Mancala game
//	@Produce		json
//	@Success		201	{object}	model.MancalaGame
//	@Failure		400	{object}	object{error string}
//	@Router			/mancala-game [post]
//	@Tags			MancalaGames
func Create(c *gin.Context) {
	game, err := entity.CreateNewMancalaGame()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": game})
}

// UpdateMancalaGame godoc
//
//	@Summary		Create new Mancala along with two players
//	@Description	Create new Mancala game
//	@Param			selected-pit	body	int	true	"Selected pit for the round"	Enums(1,2,3,4,5,6)
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	model.MancalaGame
//	@Failure		400	{object}	object{error string}
//	@Router			/mancala-game/{id} [put]
//	@Tags			MancalaGames
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	util.LogBoard(game) //DEBUG PURPOSES

	c.JSON(http.StatusOK, game)
}
