package route

import (
	_ "github.com/Abeldlp/bol-assignment/mancala-api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func AppendDocumentationRoute(r *gin.Engine) {
	swagger := r.Group("/docs/*any")

	{
		swagger.GET("", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
}
