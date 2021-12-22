package router

import (
	"go_web/api"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	apiRouter := router.Group("/api")
	{
		nameGroup := apiRouter.Group("/name")
		{
			nameGroup.POST("", api.Savename)
			nameGroup.GET("", api.GetAllNames)
			nameGroup.PUT("", api.Savename)

			nameGroup.GET("/:id", api.GetNameById)
			nameGroup.DELETE("/:id", api.DeleteNameById)
		}

	}
	return router
}
