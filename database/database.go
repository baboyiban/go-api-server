package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/baboyiban/go-api-server/models"
)

func getDSN() string {
	user := LookupEnv("DB_USER", "root")
	pass := LookupEnv("DB_PASSWORD", "password")
	host := LookupEnv("DB_HOST", "127.0.0.1")
	port := LookupEnv("DB_PORT", "3306")
	name := LookupEnv("DB_NAME", "my_database")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)
}

func LookupEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func InitDB() *gorm.DB {
	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}

	err = db.AutoMigrate(
		&models.Zone{},
		// &models.Delivery{},
		// &models.Vehicle{},
		// &models.OperationRecord{},
		// &models.OperationDelivery{},
		// &models.Employee{},
	)
	if err != nil {
		log.Fatalf("데이터베이스 마이그레이션 실패: %v", err)
	}
	log.Println("데이터베이스 연결 및 마이그레이션 성공")
	return db
}
