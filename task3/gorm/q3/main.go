package main

import (
	"fmt"
	"log"
	"my-homework/task3/gorm/model"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
题目3：钩子函数
继续使用博客系统的模型。
要求 ：
为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
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
	//为 Post 模型添加一个钩子函数，在文章创建时自动更新用户的文章数量统计字段。
	createPost(db, 1)
	//createComment(db, 44, 1)
	//为 Comment 模型添加一个钩子函数，在评论删除时检查文章的评论数量，如果评论数量为 0，则更新文章的评论状态为 "无评论"。
	deleteComment(db, 493)
}

// 模拟创建评论
func createComment(db *gorm.DB, postId, userId uint64) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic: %v", p)
		}
	}()
	comment := model.Comment{
		Content: "test2",
		PostID:  postId,
		UserID:  userId,
	}
	err = db.Create(&comment).Error
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("comment created")
	return
}
func createPost(db *gorm.DB, userId uint64) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic: %v", p)
		}
	}()
	post := model.Post{
		Title:   "test1",
		Content: "test1",
		UserID:  userId,
	}
	err = db.Create(&post).Error
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("post created")

	var user model.User
	err = db.Model(&model.User{}).Where("id = ?", userId).First(&user).Error
	if err != nil {
		log.Fatal(err)
		return
	}
	log.Println("user post num: ", user.PostNum)
	return
}
func deleteComment(db *gorm.DB, commentId uint64) (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("panic: %v", p)
		}
	}()
	err = db.Transaction(func(tx *gorm.DB) error {
		var comment model.Comment
		queryErr := tx.Model(&model.Comment{}).Where("id=?", commentId).First(&comment).Error
		if queryErr != nil {
			return queryErr
		}
		return tx.Delete(&comment).Error
	})
	return
}
