package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 参考:
// https://gorm.io/zh_CN/docs/connecting_to_the_database.html
func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3357)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// 如果 DB 不存在 ，提示如下：
	// [error] failed to initialize database, got error Error 1049: Unknown database 'dbname'
	// panic: Error 1049: Unknown database 'dbname'
	if err != nil {
		panic(err)
	}
}
