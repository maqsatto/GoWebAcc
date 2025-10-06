package routes

import (
	"accountantapp/go-service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterTransferRoutes(r *gin.Engine) {
	transfers := r.Group("/transfers")
	{
		transfers.POST("/", controllers.CreateTransfer)
		transfers.GET("/", controllers.GetAllTransfers)
		transfers.GET("/:id", controllers.GetTransferByID)
		transfers.GET("/account/:accountId", controllers.GetTransfersByAccount)
		transfers.DELETE("/:id", controllers.DeleteTransfer)
	}
}
