package handlers

import (
	"github.com/baboyiban/go-api-server/models"
	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// CreateRegion godoc
// @Summary      지역 생성
// @Description  새로운 지역을 생성합니다.
// @Tags         region
// @Accept       json
// @Produce      json
// @Param        region  body      models.Region  true  "지역 정보"
// @Success      201     {object}  models.Region
// @Failure      400     {object}  map[string]string
// @Failure      500     {object}  map[string]string
// @Router       /api/region [post]
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
	router.POST("/api/trip-log-a", CreateHandler[models.TripLogA](db))
	router.POST("/api/trip-log-b", CreateHandler[models.TripLogB](db))
	router.POST("/api/delivery_log", CreateHandler[models.DeliveryLog](db))
	router.POST("/api/employee", CreateHandler[models.Employee](db))
}
