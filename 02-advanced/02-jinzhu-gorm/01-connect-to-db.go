package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// MySQL
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()

	// SQLite3
	db, err := gorm.Open("sqlite3", "/Users/finnley/Coding/go-learning/02-advanced/03-jinzhu-gorm/gorm.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 创建表，自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserInfo{})

	// 创建数据行
	//u1 := UserInfo{
	//	ID:     1,
	//	Name:   "张三",
	//	Gender: "男",
	//	Hobby:  "长跑",
	//}
	//db.Create(&u1)

	// 查询
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%#v\n", u)

	// 更新
	//db.Model(&u).Update("hobby", "游戏")

	// 删除
	db.Delete(&u)
}
