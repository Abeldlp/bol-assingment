package route

import (
	"github.com/Abeldlp/bol-assignment/mancala-api/controller"
	"github.com/gin-gonic/gin"
)

func AppendMancalaGameRoute(r *gin.Engine) {
	mancalaGame := r.Group("/mancala-game")

	{
		mancalaGame.GET("", controller.GetAll)
		mancalaGame.GET("/:id", controller.GetById)
		mancalaGame.POST("", controller.Create)
		mancalaGame.PUT("/:id", controller.Update)
	}
}
