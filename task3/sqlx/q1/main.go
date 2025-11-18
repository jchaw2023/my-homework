package main

import (
	"log"
	"my-homework/task3/sqlx/q1/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
题目 ：使用 Go 的 sqlx 包连接 MySQL 数据库，执行一个简单的查询操作。
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

	log.Println("查询技术部员工信息")
	var employees []model.Employee
	err = db.Select(&employees, "SELECT id, name, department, salary FROM employees WHERE department = ?", "技术部")
	if err != nil {
		log.Fatal(err)
	}
	for _, employee := range employees {
		log.Println("技术部员工信息：", employee.ID, employee.Name, employee.Department, employee.Salary)
	}
	log.Println("查询薪资最高的员工信息")
	var maxSalaryEmployee model.Employee
	err = db.Get(&maxSalaryEmployee, "SELECT id, name, department,  MAX(salary) as salary FROM employees")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("薪资最高的员工信息：", maxSalaryEmployee.ID, maxSalaryEmployee.Name, maxSalaryEmployee.Department, maxSalaryEmployee.Salary)
}
