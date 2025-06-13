package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/baboyiban/go-api-server/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 전체 조회
func GetAllHandler[T any](db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var results []T
		if err := db.Find(&results).Error; err != nil {
			c.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("Failed to fetch records: %v", err)})
			return
		}
		c.JSON(200, results)
	}
}

// ID 조회
func GetByIDHandler[T any](db *gorm.DB, idField string) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		var item T

		if err := db.Where(fmt.Sprintf("%s = ?", idField), id).First(&item).Error; err != nil {
			c.AbortWithStatusJSON(404, gin.H{"error": "Record not found"})
			return
		}

		c.JSON(200, item)
	}
}

// 각 필드 조회
func GetByField[T any](db *gorm.DB, validFields map[string]bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 쿼리 파라미터 추출 (필터 조건)
		queryParams := c.Request.URL.Query()
		validParams := make(map[string]string)
		for field, values := range queryParams {
			if field == "sort" {
				continue // 정렬 파라미터는 별도 처리
			}
			if !validFields[field] {
				continue // 유효하지 않은 필드는 무시
			}
			if len(values) > 0 {
				validParams[field] = values[0]
			}
		}

		// 2. 시간 필드 파싱
		timeFields := map[string]bool{
			"completion_at":         true,
			"first_transport_time":  true,
			"input_time":            true,
			"second_transport_time": true,
			"registration_time":     true,
			"registered_at":         true,
			"saturated_at":          true,
			"start_time":            true,
			"end_time":              true,
		}
		parsedParams := make(map[string]any)
		for field, value := range validParams {
			if timeFields[field] {
				t, err := time.Parse(time.RFC3339, value)
				if err != nil {
					c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid time format for field %s", field)})
					return
				}
				parsedParams[field] = t
			} else {
				parsedParams[field] = value
			}
		}

		// 3. 쿼리 빌드 (필터 조건 적용)
		var results []T
		query := db.Model(&results)
		for field, value := range parsedParams {
			query = query.Where(fmt.Sprintf("%s = ?", field), value)
		}

		// 4. 정렬 조건 적용 (예: sort=-registered_at -> registered_at DESC)
		if sortParam := c.Query("sort"); sortParam != "" {
			sortField := sortParam
			order := "ASC"
			if sortParam[0] == '-' {
				sortField = sortParam[1:]
				order = "DESC"
			}
			if !validFields[sortField] {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid sort field: %s", sortField)})
				return
			}
			query = query.Order(fmt.Sprintf("%s %s", sortField, order))
		}

		// 5. 결과 반환
		if err := query.Find(&results).Error; err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Database query failed"})
			return
		}
		c.JSON(http.StatusOK, results)
	}
}

func RegisterReadHandlers(router *gin.Engine, db *gorm.DB) {
	// 전체 조회
	router.GET("/api/region", GetAllHandler[models.Region](db))
	router.GET("/api/package", GetAllHandler[models.Package](db))
	router.GET("/api/vehicle", GetAllHandler[models.Vehicle](db))
	router.GET("/api/trip_log", GetAllHandler[models.TripLog](db))
	router.GET("/api/delivery_log", GetAllHandler[models.DeliveryLog](db))
	router.GET("/api/employee", GetAllHandler[models.Employee](db))

	// 아이디 조회
	router.GET("/api/region/:id", GetByIDHandler[models.Region](db, "region_id"))
	router.GET("/api/package/:id", GetByIDHandler[models.Package](db, "package_id"))
	router.GET("/api/vehicle/:id", GetByIDHandler[models.Vehicle](db, "internal_id"))
	router.GET("/api/trip_log/:id", GetByIDHandler[models.TripLog](db, "trip_id"))
	router.GET("/api/delivery_log/:id", GetByIDHandler[models.DeliveryLog](db, "trip_id"))
	router.GET("/api/employee/:id", GetByIDHandler[models.Employee](db, "employee_id"))

	// 각 필드로 검색
	router.GET("/api/region/search", GetByField[models.Region](db, map[string]bool{
		"region_id":        true,
		"region_name":      true,
		"coord_x":          true,
		"coord_y":          true,
		"max_capacity":     true,
		"current_capacity": true,
		"is_full":          true,
		"saturated_at":     true,
	}))
	router.GET("/api/package/search", GetByField[models.Package](db, map[string]bool{
		"package_id":     true,
		"package_type":   true,
		"region_id":      true,
		"package_status": true,
		"registered_at":  true,
	}))
	router.GET("/api/vehicle/search", GetByField[models.Vehicle](db, map[string]bool{
		"internal_id":        true,
		"vehicle_id":         true,
		"current_load":       true,
		"max_load":           true,
		"led_status":         true,
		"needs_confirmation": true,
		"coord_x":            true,
		"coord_y":            true,
	}))
	router.GET("/api/trip_log/search", GetByField[models.TripLog](db, map[string]bool{
		"trip_id":    true,
		"vehicle_id": true,
		"start_time": true,
		"end_time":   true,
		"status":     true,
	}))
	router.GET("/api/delivery_log/search", GetByField[models.DeliveryLog](db, map[string]bool{
		"trip_id":               true,
		"package_id":            true,
		"region_id":             true,
		"load_order":            true,
		"registration_time":     true,
		"first_transport_time":  true,
		"input_time":            true,
		"second_transport_time": true,
		"completion_at":         true,
	}))
	router.GET("/api/employee/search", GetByField[models.Employee](db, map[string]bool{
		"employee_id": true,
		"password":    true,
		"position":    true,
		"is_active":   true,
	}))
}
