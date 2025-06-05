package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := initDB()
	router := gin.Default()
	api := router.Group("/api")
	{
		// Zone
		api.POST("/zones", createHandler[Zone](db))
		api.GET("/zones", getAllHandler[Zone](db))
		api.PUT("/zones/:id", updateHandler[Zone](db))
		api.DELETE("/zones/:id", deleteHandler[Zone](db))
		// Delivery
		api.POST("/deliverys", createHandler[Delivery](db))
		api.GET("/deliverys", getAllHandler[Delivery](db))
		api.PUT("/deliverys/:id", updateHandler[Delivery](db))
		api.DELETE("/deliverys/:id", deleteHandler[Delivery](db))
		// Vehicle
		api.POST("/vehicles", createHandler[Vehicle](db))
		api.GET("/vehicles", getAllHandler[Vehicle](db))
		api.PUT("/vehicles/:id", updateHandler[Vehicle](db))
		api.DELETE("/vehicles/:id", deleteHandler[Vehicle](db))
		// OperationRecord
		api.POST("/operation_records", createHandler[OperationRecord](db))
		api.GET("/operation_records", getAllHandler[OperationRecord](db))
		api.PUT("/operation_records/:id", updateHandler[OperationRecord](db))
		api.DELETE("/operation_records/:id", deleteHandler[OperationRecord](db))
		// OperationDelivery
		api.POST("/operation_Deliverys", createHandler[OperationDelivery](db))
		api.GET("/operation_Deliverys", getAllHandler[OperationDelivery](db))
		api.PUT("/operation_Deliverys/:id", updateHandler[OperationDelivery](db))
		api.DELETE("/operation_Deliverys/:id", deleteHandler[OperationDelivery](db))
	}
	port := ":" + lookupEnv("API_PORT", "")
	log.Printf("서버가 %s 포트에서 실행 중...", port)
	if err := router.Run(port); err != nil {
		log.Fatalf("서버 실행 실패: %v", err)
	}
}
