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
	getEnv := func(key, fallback string) string {
		if v, ok := os.LookupEnv(key); ok {
			return v
		}
		return fallback
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		getEnv("DB_USER", "root"),
		getEnv("DB_PASSWORD", "password"),
		getEnv("DB_HOST", "127.0.0.1"),
		getEnv("DB_PORT", "3306"),
		getEnv("DB_NAME", "my_database"),
	)
}

func InitDB() *gorm.DB {
	db, err := gorm.Open(mysql.Open(getDSN()), &gorm.Config{})
	if err != nil {
		log.Fatalf("DB 연결 실패: %v", err)
	}

	// 필요할 때만 마이그레이션 수행
	// autoMigrateAll(db)

	log.Println("DB 연결 완료")
	return db
}

// autoMigrateAll 모든 모델에 대해 자동 마이그레이션 수행
func autoMigrateAll(db *gorm.DB) {
	modelsToMigrate := []any{
		&models.Region{},
		&models.Vehicle{},
		&models.Employee{},
		&models.Package{},
		&models.TripLogA{},
		&models.TripLogB{},
		&models.DeliveryLog{},
	}

	for _, m := range modelsToMigrate {
		name := fmt.Sprintf("%T", m)
		log.Printf("마이그레이션 시작: %s", name)
		if err := db.AutoMigrate(m); err != nil {
			log.Fatalf("마이그레이션 실패: %s, 에러: %v", name, err)
		}
		log.Printf("마이그레이션 완료: %s", name)
	}
}
