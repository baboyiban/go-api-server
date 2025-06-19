package main

import (
	"log"
	"os"

	_ "github.com/baboyiban/go-api-server/docs"
	"github.com/baboyiban/go-api-server/middleware"
	"github.com/baboyiban/go-api-server/service"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"

	"github.com/baboyiban/go-api-server/database"
	"github.com/baboyiban/go-api-server/handlers"
)

// @title           Go API Server
// @version         1.0
// @description     패키지 운송 시스템 API 문서입니다.
// @host            localhost:3000
// @BasePath        /
func main() {
	// 환경변수에서 GIN_MODE 읽어서 없으면 default는 debug 모드
	mode := getEnv("GIN_MODE", gin.DebugMode)
	gin.SetMode(mode)

	db := database.InitDB()
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	registerRoutes(router, db)

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

func registerRoutes(router *gin.Engine, db *gorm.DB) {
	regionService := service.NewRegionService(db)
	regionHandler := handlers.NewRegionHandler(regionService)
	router.POST("/api/region", regionHandler.CreateRegion)
	router.GET("/api/region/:id", regionHandler.GetRegionByID)
	router.PUT("/api/region/:id", regionHandler.UpdateRegion)
	router.DELETE("/api/region/:id", regionHandler.DeleteRegion)
	router.GET("/api/region", regionHandler.ListRegions)
	router.GET("/api/region/search", regionHandler.SearchRegions)

	packageService := service.NewPackageService(db)
	packageHandler := handlers.NewPackageHandler(packageService)
	router.POST("/api/package", packageHandler.CreatePackage)
	router.GET("/api/package/:id", packageHandler.GetPackageByID)
	router.PUT("/api/package/:id", packageHandler.UpdatePackage)
	router.DELETE("/api/package/:id", packageHandler.DeletePackage)
	router.GET("/api/package", packageHandler.ListPackages)
	router.GET("/api/package/search", packageHandler.SearchPackages)

	vehicleService := service.NewVehicleService(db)
	vehicleHandler := handlers.NewVehicleHandler(vehicleService)
	router.POST("/api/vehicle", vehicleHandler.CreateVehicle)
	router.GET("/api/vehicle/:id", vehicleHandler.GetVehicleByID)
	router.PUT("/api/vehicle/:id", vehicleHandler.UpdateVehicle)
	router.DELETE("/api/vehicle/:id", vehicleHandler.DeleteVehicle)
	router.GET("/api/vehicle", vehicleHandler.ListVehicles)
	router.GET("/api/vehicle/search", vehicleHandler.SearchVehicles)

	tripLogService := service.NewTripLogService(db)
	tripLogHandler := handlers.NewTripLogHandler(tripLogService)
	router.POST("/api/trip-log", tripLogHandler.CreateTripLog)
	router.GET("/api/trip-log/:id", tripLogHandler.GetTripLogByID)
	router.PUT("/api/trip-log/:id", tripLogHandler.UpdateTripLog)
	router.DELETE("/api/trip-log/:id", tripLogHandler.DeleteTripLog)
	router.GET("/api/trip-log", tripLogHandler.ListTripLogs)
	router.GET("/api/trip-log/search", tripLogHandler.SearchTripLogs)

	tripLogBService := service.NewTripLogBService(db)
	tripLogBHandler := handlers.NewTripLogBHandler(tripLogBService)
	router.POST("/api/trip-log-b", tripLogBHandler.CreateTripLogB)
	router.GET("/api/trip-log-b/:id", tripLogBHandler.GetTripLogBByID)
	router.PUT("/api/trip-log-b/:id", tripLogBHandler.UpdateTripLogB)
	router.DELETE("/api/trip-log-b/:id", tripLogBHandler.DeleteTripLogB)
	router.GET("/api/trip-log-b", tripLogBHandler.ListTripLogBs)
	router.GET("/api/trip-log-b/search", tripLogBHandler.SearchTripLogBs)

	deliveryLogService := service.NewDeliveryLogService(db)
	deliveryLogHandler := handlers.NewDeliveryLogHandler(deliveryLogService)
	router.POST("/api/delivery-log", deliveryLogHandler.CreateDeliveryLog)
	router.GET("/api/delivery-log/:id", deliveryLogHandler.GetDeliveryLogByID)
	router.PUT("/api/delivery-log/:id", deliveryLogHandler.UpdateDeliveryLog)
	router.DELETE("/api/delivery-log/:id", deliveryLogHandler.DeleteDeliveryLog)
	router.GET("/api/delivery-log", deliveryLogHandler.ListDeliveryLogs)
	router.GET("/api/delivery-log/search", deliveryLogHandler.SearchDeliveryLogs)

	employeeService := service.NewEmployeeService(db)
	employeeHandler := handlers.NewEmployeeHandler(employeeService)
	router.POST("/api/employee", middleware.AuthRequired("관리직"), employeeHandler.CreateEmployee)
	router.GET("/api/employee/:id", middleware.AuthRequired("관리직"), employeeHandler.GetEmployeeByID)
	router.PUT("/api/employee/:id", middleware.AuthRequired("관리직"), employeeHandler.UpdateEmployee)
	router.DELETE("/api/employee/:id", middleware.AuthRequired("관리직"), employeeHandler.DeleteEmployee)
	router.GET("/api/employee", middleware.AuthRequired("관리직"), employeeHandler.ListEmployees)
	router.GET("/api/employee/search", middleware.AuthRequired("관리직"), employeeHandler.SearchEmployees)

	// auth
	authService := service.NewAuthService(db)
	authHandler := handlers.NewAuthHandler(authService)
	router.POST("/api/auth/login", authHandler.Login)
}
