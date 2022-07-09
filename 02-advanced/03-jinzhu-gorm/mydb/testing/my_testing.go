package main

import (
	"context"
	"database/sql"
	"fmt"
	"go-learning/02-advanced/03-jinzhu-gorm/mydb"
	"log"
	"strings"
)

func init() {
	log.Println("register mydb driver")
	sql.Register("mydb", &mydb.Driver{})
}

type MyUser struct {
	Id      int
	Name    string
	Age     int
	Version string
}

//func TestDb(t *testing.T) {
func main() {
	//db, err := sql.Open("mydb", "root:root@tcp(127.0.0.1:3357)/gorm_test")
	db, err := sql.Open("mydb", "@tcp(127.0.0.1:3358)/?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("some error %s", err.Error())
		//t.Errorf("some error %s", err.Error())
	}
	//err = db.Ping()
	err = db.PingContext(context.Background())

	fmt.Printf("has err: %v\n", err)
	if err != nil && strings.Contains(err.Error(), "mysql") {
		fmt.Println("mysql service discovery")
	}
	//fmt.Println(db)
	//rows, err := db.Query("select name,age,version from demoapp")
	//if err != nil {
	//	log.Fatal("some wrong for query", err.Error())
	//}
	//for rows.Next() {
	//	var user MyUser
	//	if err := rows.Scan(&user.Name, &user.Age, &user.Version); err != nil {
	//		log.Println("scan value erro", err.Error())
	//	} else {
	//		log.Println(user)
	//	}
	//}
}
