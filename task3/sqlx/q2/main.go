package main

import (
	"log"
	"my-homework/task3/sqlx/q2/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
题目 ：使用 Go 的 sqlx 包连接 MySQL 数据库，执行一个简单的插入操作。
考察点 ：sqlx 包的使用、MySQL 数据库的连接。
*/
func main() {
	db, err := sqlx.Connect("mysql", "root:123qwe@tcp(127.0.0.1:3306)/my_test_db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	var books []model.Book
	err = db.Select(&books, "SELECT * FROM books WHERE price > ?", "50")
	if err != nil {
		log.Fatal(err)
	}
	for _, book := range books {
		log.Println("图书信息：", book.ID, book.Title, book.Author, book.Price)
	}
}
