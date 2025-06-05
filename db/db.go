package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDSN() string {
	user := lookupEnv("DB_USER", "root")
	pass := lookupEnv("DB_PASSWORD", "password")
	host := lookupEnv("DB_HOST", "127.0.0.1")
	port := lookupEnv("DB_PORT", "3306")
	name := lookupEnv("DB_NAME", "my_database")
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)
}

func lookupEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func initDB() *gorm.DB {
	dsn := getDSN()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}

	err = db.AutoMigrate(
		&Delivery{},
		&Zone{},
		&Vehicle{},
		&OperationRecord{},
		&OperationProduct{},
	)
	if err != nil {
		log.Fatalf("데이터베이스 마이그레이션 실패: %v", err)
	}
	log.Println("데이터베이스 연결 및 마이그레이션 성공")
	return db
}
