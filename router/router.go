package router

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	router := gin.Default()

	user := router.Group("user")
	{
		user.POST("register")
		user.POST("login")
	}

	router.Run(":8080")
	return router
}
