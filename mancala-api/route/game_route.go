package route

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/controller"
	"github.com/gin-gonic/gin"
)

func AppendMancalaGameRoute(r *gin.Engine) {
	mancalaGame := r.Group("/mancala-game")

	{
		mancalaGame.GET("", controller.GetAllMancalaGames)
		mancalaGame.GET("/:id", controller.GetMancalaGame)
		mancalaGame.POST("", controller.CreateMancalaGame)
		mancalaGame.PUT("/:id", controller.UpdateMancalaGame)
	}
}
