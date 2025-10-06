package routes

import (
	"accountantapp/go-service/internal/controllers"
	"accountantapp/go-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterAccountRoutes(r *gin.Engine) {
	accounts := r.Group("/accounts")
	accounts.Use(middleware.AuthMiddleware())
	{
		accounts.POST("/", controllers.CreateAccount)
		accounts.GET("/", controllers.GetAllAccounts)
		accounts.GET("/:id", middleware.ValidateAccountID(), controllers.GetAccountByID)
		accounts.PUT("/:id", middleware.ValidateAccountID(), controllers.UpdateAccount)
		accounts.DELETE("/:id", middleware.ValidateAccountID(), controllers.DeleteAccount)
	}
}
