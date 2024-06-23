package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hana-boy/golang_blog/models"
)

func main() {
	// Ginエンジンのインスタンスを作成
	r := gin.Default()

	// タスクを取得するエンドポイント
	r.GET("/tasks", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetTasks())
	})

	// 新しいタスクを作成するエンドポイント
	r.POST("/tasks", func(c *gin.Context) {
		task, err := models.CreateTask(c)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, task)
	})

	// タスクを更新するエンドポイント
	r.PUT("/tasks/:id", func(c *gin.Context) {
		task, err := models.UpdateTask(c)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, task)
	})

	// タスクを削除するエンドポイント
	r.DELETE("/tasks/:id", func(c *gin.Context) {
		err := models.DeleteTask(c)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Task deleted"})
	})

	// 記事の一覧を取得するエンドポイント
	r.GET("/articles", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetArticleIndex())
	})

	// 記事の詳細を取得するエンドポイント
	r.GET("/articles/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.GetArticleDetail(c))
	})

	// 新しい記事を作成するエンドポイント
	r.POST("/articles", func(c *gin.Context) {
		task, err := models.CreateArticle(c)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, task)
	})

	// 記事を更新するエンドポイント
	r.PUT("/articles/:id", func(c *gin.Context) {
		task, err := models.UpdateArticle(c)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, task)
	})

	// 記事を削除するエンドポイント
	r.DELETE("/articles/:id", func(c *gin.Context) {
		err := models.DeleteArticle(c)
		if err != nil {
			log.Fatalln(err)
		}
		c.JSON(http.StatusOK, gin.H{"message": "Article deleted"})
	})

	// 8080ポートでサーバーを起動
	r.Run(":8080")
}
