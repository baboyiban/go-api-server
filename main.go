package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/baboyiban/go-api-server/database"
	"github.com/baboyiban/go-api-server/handlers"
)

func main() {
	// 환경변수에서 GIN_MODE 읽어서 없으면 default는 debug 모드
	mode := getEnv("GIN_MODE", gin.DebugMode)
	gin.SetMode(mode)

	db := database.InitDB()
	router := gin.Default()

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
	handlers.RegisterCreateHandlers(router, db)
	handlers.RegisterReadHandlers(router, db)
	handlers.RegisterDeleteHandlers(router, db)
	handlers.RegisterUpdateHandlers(router, db)
}
