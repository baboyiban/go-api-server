package handlers

import (
    "fmt"

    "github.com/gin-gonic/gin"
    "gorm.io/gorm"
)

func createHandler[T any](db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var item T
        if err := c.ShouldBindJSON(&item); err != nil {
            c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
            return
        }
        result := db.Create(&item)
        if result.Error != nil {
            c.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("생성 실패: %v", result.Error.Error())})
            return
        }
        c.JSON(201, item)
    }
}

func getAllHandler[T any](db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        var items []T
        result := db.Find(&items)
        if result.Error != nil {
            c.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("조회 실패: %v", result.Error.Error())})
            return
        }
        c.JSON(200, items)
    }
}

func updateHandler[T any](db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var item T
        var existingItem T
        result := db.Where("ID = ?", id).First(&existingItem)
        if result.Error != nil {
            if result.Error == gorm.ErrRecordNotFound {
                c.AbortWithStatusJSON(404, gin.H{"error": "리소스를 찾을 수 없습니다."})
            } else {
                c.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("리서스 조회 실패: %v", result.Error.Error())})
            }
            return
        }
        if err := c.ShouldBindJSON(&item); err != nil {
            c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
            return
        }
        result = db.Model(&existingItem).Updates(item)
        if result.Error != nil {
            c.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("업데이트 실패: %v", result.Error.Error())})
            return
        }
        c.JSON(200, existingItem)
    }
}

func deleteHandler[T any](db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        id := c.Param("id")
        var item T
        result := db.Where("ID = ?", id).Delete(&item)
        if result.Error != nil {
            c.AbortWithStatusJSON(500, gin.H{"error": fmt.Sprintf("삭제 실패: %v", result.Error.Error())})
            return
        }
        if result.RowsAffected == 0 {
            c.AbortWithStatusJSON(404, gin.H{"error": "리소스를 찾을 수 없거나 이미 삭제되었습니다."})
            return
        }
        c.Status(204)
    }
}
