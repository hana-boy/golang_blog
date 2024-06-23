package models

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Task struct {
	ID          uint       `gorm:"primary_key"`
	Task        string     `gorm:"size:255"`
	IsCompleted bool       `gorm:"default:false"`
	CreatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt   *time.Time `gorm:"index"`
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
