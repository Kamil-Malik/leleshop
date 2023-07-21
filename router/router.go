package router

import (
	"leleshop/controller"
	"leleshop/middleware"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	user := router.Group("user")
	{
		user.POST("register", controller.Register)
		user.POST("login", controller.Login)
		user.GET("profile/:user_name", middleware.Authentication(), controller.GetUserProfile)
	}

	router.Run(":8080")
	return router
}
