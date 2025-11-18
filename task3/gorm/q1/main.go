package main

import (
	"log"
	"my-homework/task3/gorm/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
要求 ：
使用Gorm定义 User 、 Post 和 Comment 模型，其中 User 与 Post 是一对多关系（一个用户可以发布多篇文章）， Post 与 Comment 也是一对多关系（一篇文章可以有多个评论）。
编写Go代码，使用Gorm创建这些模型对应的数据库表。
*/
func main() {
	db, err := gorm.Open(mysql.Open("root:123qwe@tcp(127.0.0.1:3306)/my_test_db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}
		sqlDB.Close()
	}()
	db.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})

}
