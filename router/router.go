package router

import (
	"leleshop/controller"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	user := router.Group("user")
	{
		user.POST("register", controller.Register)
		user.POST("login", controller.Login)
	}

	router.Run(":8080")
	return router
}
