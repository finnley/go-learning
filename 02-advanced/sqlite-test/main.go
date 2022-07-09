package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main02() {
	fmt.Println("==> Hello,World!")
}

type Product struct {
	gorm.Model
	Title string
	Code  string
	Price uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.orm_test"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// 插入内容
	db.Create(&Product{Title: "新款手机", Code: "D42", Price: 1000})
	db.Create(&Product{Title: "新款电脑", Code: "D43", Price: 3500})

	// 读取内容
	var product Product
	//orm_test.First(&product, 1)                  // find product with integer primary key
	db.First(&product, "code = ?", "D424") // find product with code D42
	//
	//// 更新操作：更新单个字段
	//orm_test.Model(&product).Update("Price", 2000)
	//
	//// 更新操作：更新多个字段
	//orm_test.Model(&product).Updates(Product{Price: 2000, Code: "F42"}) // non-zero fields
	//orm_test.Model(&product).Updates(map[string]interface{}{"Price": 2000, "Code": "F42"})
	//
	//// 删除操作：
	//orm_test.Delete(&product, 1)
}
