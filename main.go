package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/baboyiban/go-api-server/database"
	"github.com/baboyiban/go-api-server/handlers"
	"github.com/baboyiban/go-api-server/models"
)

func main() {
	db := database.InitDB()
	router := gin.Default()
	api := router.Group("/api")

	registerRoutes(api, db)

	port := getEnv("API_PORT", "8080")
	addr := ":" + port
	log.Printf("서버가 %s 포트에서 실행 중...", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}

func getEnv(key, fallback string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	return fallback
}

func registerRoutes(api *gin.RouterGroup, db *gorm.DB) {

	// Region
	api.POST("/region", handlers.CreateHandler[models.Region](db))
	api.GET("/region", handlers.GetAllHandler[models.Region](db))
	api.PUT("/region/:id", handlers.UpdateHandler[models.Region](db))
	api.DELETE("/region/:id", handlers.DeleteHandler[models.Region](db))

	// Package
	api.POST("/package", handlers.CreateHandler[models.Package](db))
	api.GET("/package", handlers.GetAllHandler[models.Package](db))
	api.PUT("/package/:id", handlers.UpdateHandler[models.Package](db))
	api.DELETE("/package/:id", handlers.DeleteHandler[models.Package](db))

	// Vehicle
	api.POST("/vehicle", handlers.CreateHandler[models.Vehicle](db))
	api.GET("/vehicle", handlers.GetAllHandler[models.Vehicle](db))
	api.PUT("/vehicle/:id", handlers.UpdateHandler[models.Vehicle](db))
	api.DELETE("/vehicle/:id", handlers.DeleteHandler[models.Vehicle](db))

	// TripLog
	api.POST("/trip_log", handlers.CreateHandler[models.TripLog](db))
	api.GET("/trip_log", handlers.GetAllHandler[models.TripLog](db))
	api.PUT("/trip_log/:id", handlers.UpdateHandler[models.TripLog](db))
	api.DELETE("/trip_log/:id", handlers.DeleteHandler[models.TripLog](db))

	// DeliveryLog
	api.POST("/delivery_log", handlers.CreateHandler[models.DeliveryLog](db))
	api.GET("/delivery_log", handlers.GetAllHandler[models.DeliveryLog](db))
	api.PUT("/delivery_log/:id", handlers.UpdateHandler[models.DeliveryLog](db))
	api.DELETE("/delivery_log/:id", handlers.DeleteHandler[models.DeliveryLog](db))
}
