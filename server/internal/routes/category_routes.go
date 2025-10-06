package routes

import (
	"accountantapp/go-service/internal/controllers"
	"accountantapp/go-service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterCategoryRoutes(r *gin.Engine) {
	cats := r.Group("/categories")
	cats.Use(middleware.AuthMiddleware())
	{
		cats.POST("/", controllers.CreateCategory)
		cats.GET("/", controllers.GetAllCategories)
		cats.GET("/:id", middleware.ValidateCategoryID(), controllers.GetCategoryByID)
		cats.PUT("/:id", middleware.ValidateCategoryID(), controllers.UpdateCategory)
		cats.DELETE("/:id", middleware.ValidateCategoryID(), controllers.DeleteCategory)
	}
}
