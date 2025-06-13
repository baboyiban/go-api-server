package handlers

import (
	"github.com/baboyiban/go-api-server/models"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
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
			// MySQL 중복 키 오류 (Error 1062) 처리
			if isDuplicateError(err) {
				c.AbortWithStatusJSON(400, gin.H{
					"error": "중복된 데이터입니다",
				})
				return
			}
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to create record"})
			return
		}

		c.JSON(201, item)
	}
}

// isDuplicateError는 MySQL 중복 키 오류(1062)를 확인하는 헬퍼 함수
func isDuplicateError(err error) bool {
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		return mysqlErr.Number == 1062 // MySQL Error 1062: Duplicate entry
	}
	return false
}

func RegisterCreateHandlers(router *gin.Engine, db *gorm.DB) {
	router.POST("/api/region", CreateHandler[models.Region](db))
	router.POST("/api/package", CreateHandler[models.Package](db))
	router.POST("/api/vehicle", CreateHandler[models.Vehicle](db))
	router.POST("/api/trip_log", CreateHandler[models.TripLog](db))
	router.POST("/api/delivery_log", CreateHandler[models.DeliveryLog](db))
	router.POST("/api/employee", CreateHandler[models.Employee](db))
}
