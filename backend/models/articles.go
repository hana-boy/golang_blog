package models

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Article struct {
	ID        uint `gorm:"primary_key"`
	UserID    int
	Title     string `gorm:"size:255"`
	Content   string
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt *time.Time `gorm:"index"`
}

// 記事の一覧を取得する関数
func GetArticleIndex() []Article {
	var articles []Article
	db := ConnectDb()
	db.Select("id, title").Find(&articles)
	return articles
}

// 記事の詳細を取得する関数
func GetArticleDetail(c *gin.Context) Article {
	var article Article
	db := ConnectDb()
	db.Where("id = ?", c.Param("id")).First(&article)
	return article
}

// 新しい記事を作成する関数
func CreateArticle(c *gin.Context) (article Article, err error) {
	db := ConnectDb()

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return article, err
	}
	db.Create(&article)
	return article, nil
}

// 記事を更新する関数
func UpdateArticle(c *gin.Context) (article Article, err error) {
	id := c.Param("id")
	db := ConnectDb()

	if err := db.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return article, err
	}

	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return article, err
	}

	db.Save(&article)
	return article, nil
}

// 記事を削除する関数
func DeleteArticle(c *gin.Context) (err error) {
	var article Article
	id := c.Param("id")
	db := ConnectDb()

	if err := db.First(&article, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Article not found"})
		return err
	}

	db.Delete(&article)
	return nil
}
