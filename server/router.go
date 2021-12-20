package server

import (
	"go_web/api"

	"github.com/gin-gonic/gin"
)

func newRouter() *gin.Engine {
	router := gin.New()

	apiRouter := router.Group("/api")
	{
		nameGroup := apiRouter.Group("/name")
		{
			nameGroup.POST("", api.Savename)
			nameGroup.GET("", api.GetAllNames)
			nameGroup.PUT("", api.Savename)

			nameGroup.GET("", api.GetNameById)
			nameGroup.DELETE("", api.DeleteNameById)
		}

	}
	return router
}
