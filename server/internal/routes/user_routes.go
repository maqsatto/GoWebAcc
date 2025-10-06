package routes

import (
	"accountantapp/go-service/internal/controllers"
	"accountantapp/go-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine) {
	users := r.Group("/users")
	{
		users.POST("/register", controllers.Register)
		users.POST("/login", controllers.Login)
		users.GET("/profile", middleware.AuthMiddleware(), controllers.GetProfile)

		// Защищенные routes
		protected := users.Group("")
		protected.Use(middleware.AuthMiddleware())
		{
			protected.GET("/", controllers.GetAllUsers)
			protected.GET("/:id", controllers.GetUserByID)
			protected.PUT("/:id", controllers.UpdateUser)
			protected.DELETE("/:id", controllers.DeleteUser)
		}
	}
}
