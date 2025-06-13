package handlers

import (
	"fmt"

	"github.com/baboyiban/go-api-server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Update 핸들러 공통 함수
func UpdateHandler[T any](db *gorm.DB, idField string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. ID로 기존 레코드 조회
		id := c.Param("id")
		var item T
		query := fmt.Sprintf("%s = ?", idField)
		if err := db.Where(query, id).First(&item).Error; err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": "Record not found"})
			return
		}

		// 2. 요청 데이터 바인딩
		if err := c.ShouldBindJSON(&item); err != nil {
			c.AbortWithStatusJSON(400, gin.H{"error": "Invalid request data"})
			return
		}

		// 3. 업데이트 실행
		if err := db.Save(&item).Error; err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to update record"})
			return
		}

		c.JSON(200, item)
	}
}

// 라우트 등록용 함수
func RegisterUpdateHandlers(router *gin.Engine, db *gorm.DB) {
	// 모델별 ID 필드명 지정
	router.PUT("/api/region/:id", UpdateHandler[models.Region](db, "region_id"))
	router.PUT("/api/package/:id", UpdateHandler[models.Package](db, "package_id"))
	router.PUT("/api/vehicle/:id", UpdateHandler[models.Vehicle](db, "internal_id"))
	router.PUT("/api/trip_log/:id", UpdateHandler[models.TripLog](db, "trip_id"))
	router.PUT("/api/delivery_log/:id", UpdateHandler[models.DeliveryLog](db, "trip_id"))
	router.PUT("/api/employee/:id", UpdateHandler[models.Employee](db, "employee_id"))
}
