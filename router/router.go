package router

import (
	"backend-service/handler"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		// v1.Use(middleware.AuthMiddleware())

		v1.GET("/", handler.HelloWorldHandler)
		v1.GET("/auth/google/login", handler.GoogleLoginHandler)
		v1.GET("/auth/google/callback", handler.GoogleCallbackHandler)
	}
	return router
}
