package handlers

import (
	"github.com/baboyiban/go-api-server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateHandler[T any](db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var item T
		if err := c.ShouldBindJSON(&item); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request data"})
			return
		}

		if err := db.Create(&item).Error; err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to create record"})
			return
		}

		c.JSON(201, item)
	}
}

func RegisterCreateHandlers(router *gin.Engine, db *gorm.DB) {
	router.POST("/api/region", CreateHandler[models.Region](db))
	router.POST("/api/package", CreateHandler[models.Package](db))
	router.POST("/api/vehicle", CreateHandler[models.Vehicle](db))
	router.POST("/api/trip_log", CreateHandler[models.TripLog](db))
	router.POST("/api/delivery_log", CreateHandler[models.DeliveryLog](db))
	router.POST("/api/employee", CreateHandler[models.Employee](db))
}
