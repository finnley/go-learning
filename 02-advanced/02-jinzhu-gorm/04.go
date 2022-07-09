package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Userinfo struct {
	Id   int
	Name string
}

func main() {
	db, err := gorm.Open("sqlite3", "/Users/finnley/go/src/go-learning/02-advanced/03-jinzhu-gorm/gorm.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SingularTable(true)

	// 创建表，自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&Userinfo{})

	// 创建表
	//db.Table("user_info").CreateTable(&Userinfo{})

	var a Userinfo
	db.First(&a)
	fmt.Println(a.Name)
}
