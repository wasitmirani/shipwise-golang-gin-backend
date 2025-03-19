package routes

import (
	"shipwise/internal/app/controllers/auth"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	apiRoutes := router.Group("/api")
	{
		authRoutes := apiRoutes.Group("/auth")
		{
			authRoutes.POST("/login", auth.Login)
			authRoutes.POST("/register", gin.WrapF(auth.Register))
		}
	}

	// userRoutes := router.Group("/user")
	// {
	// 	userRoutes.GET("/:id", user.GetUser)
	// 	userRoutes.PUT("/:id", user.UpdateUser)
	// }
}
