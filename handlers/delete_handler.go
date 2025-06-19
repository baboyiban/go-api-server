package handlers

import (
	"fmt"

	"github.com/baboyiban/go-api-server/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Delete 핸들러 공통 함수
func DeleteHandler[T any](db *gorm.DB, idField string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var item T

		// 동적 쿼리 생성 (예: "vehicle_id = ?")
		query := fmt.Sprintf("%s = ?", idField)
		result := db.Where(query, id).Delete(&item)

		if result.Error != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": "Failed to delete record"})
			return
		}

		if result.RowsAffected == 0 {
			c.AbortWithStatusJSON(404, gin.H{"error": "Record not found"})
			return
		}

		c.Status(204) // No Content
	}
}

// 라우트 등록용 함수
func RegisterDeleteHandlers(router *gin.Engine, db *gorm.DB) {
	// 모델별 ID 필드명 지정
	router.DELETE("/api/region/:id", DeleteHandler[models.Region](db, "region_id"))
	router.DELETE("/api/package/:id", DeleteHandler[models.Package](db, "package_id"))
	router.DELETE("/api/vehicle/:id", DeleteHandler[models.Vehicle](db, "interal_id"))
	router.DELETE("/api/trip-log-a/:id", DeleteHandler[models.TripLogA](db, "trip_id"))
	router.DELETE("/api/trip-log-b/:id", DeleteHandler[models.TripLogB](db, "trip_id"))
	router.DELETE("/api/delivery_log/:id", DeleteHandler[models.DeliveryLog](db, "trip_id"))
	router.DELETE("/api/employee/:id", DeleteHandler[models.Employee](db, "employee_id"))
}
