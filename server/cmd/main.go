package main

import (
	"accountantapp/go-service/internal/config"
	"accountantapp/go-service/internal/database"
	"accountantapp/go-service/internal/logger"
	"accountantapp/go-service/internal/middleware"
	"accountantapp/go-service/internal/routes"
	"accountantapp/go-service/internal/validators"
	"time"

	"github.com/gin-gonic/gin"
)

func healthCheck(c *gin.Context) {
	dbStatus := "OK"
	if err := database.DB.Exec("SELECT 1").Error; err != nil {
		dbStatus = "ERROR"
	}

	c.JSON(200, gin.H{
		"status":    "OK",
		"timestamp": time.Now().Format(time.RFC3339),
		"database":  dbStatus,
		"version":   "1.0.0",
	})
}

func main() {
	// Инициализация
	config.Load()
	database.Connect()
	validators.Init()
	logger.Init()

	// Режим Gin
	gin.SetMode(gin.ReleaseMode)

	r := gin.Default()

	// Middleware
	r.Use(middleware.CORSMiddleware())
	r.Use(logger.GinLogger())
	r.Use(gin.Recovery())
	r.Use(middleware.RateLimit(100, time.Minute)) // 100 запросов в минуту

	// Routes
	routes.RegisterUserRoutes(r)
	routes.RegisterAccountRoutes(r)
	routes.RegisterTransactionRoutes(r)
	routes.RegisterCategoryRoutes(r)
	routes.RegisterTemplateRoutes(r)
	routes.RegisterSettingsRoutes(r)
	routes.RegisterTransferRoutes(r)

	// Health check
	r.GET("/health", healthCheck)

	logger.Log.Info("Server starting on :8080")
	r.Run(":8080")
}
