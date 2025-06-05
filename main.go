package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/baboyiban/go-api-server/database"
	"github.com/baboyiban/go-api-server/handlers"
	"github.com/baboyiban/go-api-server/models"
)

func main() {
	db := database.InitDB()
	router := gin.Default()
	api := router.Group("/api")
	{
		// Zone
		api.POST("/zones", handlers.CreateHandler[models.Zone](db))
		api.GET("/zones", handlers.GetAllHandler[models.Zone](db))
		api.PUT("/zones/:id", handlers.UpdateHandler[models.Zone](db))
		api.DELETE("/zones/:id", handlers.DeleteHandler[models.Zone](db))
		// Delivery
		api.POST("/deliverys", handlers.CreateHandler[models.Delivery](db))
		api.GET("/deliverys", handlers.GetAllHandler[models.Delivery](db))
		api.PUT("/deliverys/:id", handlers.UpdateHandler[models.Delivery](db))
		api.DELETE("/deliverys/:id", handlers.DeleteHandler[models.Delivery](db))
		// Vehicle
		api.POST("/vehicles", handlers.CreateHandler[models.Vehicle](db))
		api.GET("/vehicles", handlers.GetAllHandler[models.Vehicle](db))
		api.PUT("/vehicles/:id", handlers.UpdateHandler[models.Vehicle](db))
		api.DELETE("/vehicles/:id", handlers.DeleteHandler[models.Vehicle](db))
		// OperationRecord
		api.POST("/operation_records", handlers.CreateHandler[models.OperationRecord](db))
		api.GET("/operation_records", handlers.GetAllHandler[models.OperationRecord](db))
		api.PUT("/operation_records/:id", handlers.UpdateHandler[models.OperationRecord](db))
		api.DELETE("/operation_records/:id", handlers.DeleteHandler[models.OperationRecord](db))
		// OperationDelivery
		api.POST("/operation_Deliverys", handlers.CreateHandler[models.OperationDelivery](db))
		api.GET("/operation_Deliverys", handlers.GetAllHandler[models.OperationDelivery](db))
		api.PUT("/operation_Deliverys/:id", handlers.UpdateHandler[models.OperationDelivery](db))
		api.DELETE("/operation_Deliverys/:id", handlers.DeleteHandler[models.OperationDelivery](db))
	}

	port := ":" + database.LookupEnv("API_PORT", "")
	log.Printf("서버가 %s 포트에서 실행 중...", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}
