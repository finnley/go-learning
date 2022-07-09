package main

import (
	"database/sql"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type User struct {
	ID int64
	//Name string `gorm:"default:'你是谁'"`
	//Name *string `gorm:"default:'你是谁'"`
	Name sql.NullString `gorm:"default:'你是谁'"`
	Age  int64
}

func main() {
	// MySQL
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()

	// SQLite3
	db, err := gorm.Open("sqlite3", "/Users/finnley/Coding/go-learning/02_advanced/03-jinzhu-gorm/gorm.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	//db.SingularTable(true)

	// 创建表，自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&User{})

	// 创建记录
	u := User{
		//Name: "李四",
		//Name: new(string),
		Name: sql.NullString{"", true},
		Age:  20,
	}

	fmt.Println(db.NewRecord(&u)) // 判断主键是否为空 true，也就是添加的记录是不存在，是新纪录
	db.Debug().Create(&u)
	fmt.Println(db.NewRecord(&u))

}
