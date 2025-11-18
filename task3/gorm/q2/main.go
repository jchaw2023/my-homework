package main

import (
	"log"
	"my-homework/task3/gorm/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
基于上述博客系统的模型定义。
要求 ：
编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息。
编写Go代码，使用Gorm查询评论数量最多的文章信息。
*/
func main() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)
	db, err := gorm.Open(mysql.Open("root:123qwe@tcp(127.0.0.1:3306)/my_test_db?parseTime=True&loc=Local"), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("1.编写Go代码，使用Gorm查询某个用户发布的所有文章及其对应的评论信息")
	var posts []model.Post
	err = db.Model(&model.Post{}).Where("user_id = ?", 1).Preload("Comments").Find(&posts).Error
	if err != nil {
		log.Fatal(err)
	} else {
		for _, post := range posts {
			log.Println("post-id: ", post.ID)
			log.Println("user: ", post.UserID)
			log.Println("post-title: ", post.Title)
			log.Println("post-content: ", post.Content)
			log.Println("post-created-at: ", post.CreatedAt)
			for _, comment := range post.Comments {
				log.Println("comment-id: ", comment.ID)
				log.Println("comment-content: ", comment.Content)
				log.Println("comment-created-at: ", comment.CreatedAt)
			}
			log.Println("-------------------------------")
		}
	}

	log.Println("2.编写Go代码，使用Gorm查询评论数量最多的文章信息")
	var post model.Post
	err = db.Model(&model.Post{}).
		Select("posts.*").
		Joins("LEFT JOIN comments ON comments.post_id = posts.id").
		Group("posts.id").
		Order("COUNT(comments.id) DESC").
		First(&post).Error
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("post-id: ", post.ID)
		log.Println("user-id: ", post.UserID)
		log.Println("post-title: ", post.Title)
		log.Println("post-content: ", post.Content)
		log.Println("post-created-at: ", post.CreatedAt)
		log.Println("--------------------------------")
	}
}
