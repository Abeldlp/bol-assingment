package route

import (
	"github.com/Abeldlp/bol-assignment/user-api/controller"
	"github.com/gin-gonic/gin"
)

func AppendUserRoute(r *gin.Engine) {
	mancalaUser := r.Group("/users")

	{
		mancalaUser.GET("/:id", controller.GetMancalaUserById())
		mancalaUser.POST("", controller.CreateMancalaUser())
	}
}
