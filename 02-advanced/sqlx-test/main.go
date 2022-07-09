package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func initDB() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3357)/gin_mysql_test?charset=utf8mb4&parseTime=True"
	// 也可以使用 MustConnect 链接不成功就panic
	db, err = sqlx.Connect("mysql", dsn) // Open+Ping
	if err != nil {
		fmt.Printf("connect DB failed, err: %v\n", err)
		return
	}
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(10)
	return
}

type user struct {
	ID   int    `orm_test:"id"`
	Age  int    `orm_test:"age"`
	Name string `orm_test:"name"`
}

// 查询单挑数据
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	err := db.Get(&u, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err: %s\n", err)
		return
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", u.ID, u.Name, u.Age)
}

// 查询多条数据
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id > ?"
	var users []user
	err := db.Select(&users, sqlStr, 0)
	if err != nil {
		fmt.Printf("query failed, err: %s\n", err)
		return
	}
	fmt.Printf("users: %#v\n", users)
}

func main() {
	if err := initDB(); err != nil {
		fmt.Printf("init DB failed, err: %v\n", err)
		return
	}
	fmt.Println("init DB success...")
	queryRowDemo()
	queryMultiRowDemo()
}
