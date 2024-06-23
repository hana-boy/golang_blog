package models

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Task struct {
	ID          uint       `gorm:"primary_key"`
	Task        string     `gorm:"size:255"`
	IsCompleted bool       `gorm:"default:false"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `gorm:"index"`
}

func ConnectDb() gorm.DB {
	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// 環境変数から接続情報を取得
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")
	dbHost := "localhost" // または環境変数から取得
	dbPort := "5432"      // または環境変数から取得

	// DSNを構築
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Tokyo",
		dbHost,
		dbUser,
		dbPassword,
		dbName,
		dbPort,
	)

	// GORMでデータベースに接続
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// データベースにテーブルを作成
	db.AutoMigrate(&Task{})
	return *db
}

// タスクを取得する関数
func GetTasks() []Task {
	var tasks []Task
	db := ConnectDb()
	db.Find(&tasks)
	return tasks
}

// 新しいタスクを作成する関数
func CreateTask(c *gin.Context) (task Task, err error) {
	db := ConnectDb()

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return task, err
	}
	db.Create(&task)
	return task, nil
}

// タスクを更新する関数
func UpdateTask(c *gin.Context) (task Task, err error) {
	id := c.Param("id")
	db := ConnectDb()

	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return task, err
	}

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return task, err
	}

	db.Save(&task)
	return task, nil
}

// タスクを削除する関数
func DeleteTask(c *gin.Context) (err error) {
	var task Task
	id := c.Param("id")
	db := ConnectDb()

	if err := db.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return err
	}

	db.Delete(&task)
	return nil
}
