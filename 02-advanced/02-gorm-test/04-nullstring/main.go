package main

import (
	"database/sql"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	//Code string
	Code  sql.NullString
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3357)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 设置全局的 logger，这个 logger 在执行每个 sql 的时候会打印每一行 sql
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold: time.Second, // 慢 SQL 阈值
			LogLevel:      logger.Info, // 日志级别
			//IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful: true, // 禁用彩色打印
		},
	)

	// 全局模式
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:            newLogger,
		AllowGlobalUpdate: true,
	})
	if err != nil {
		panic(err)
	}

	// 定义一个表结构，将表结构直接生成对应的表
	// 迁移 schema
	_ = db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//db.Create(&Product{Code: sql.NullString{"D42", true}, Price: 100})

	// Read
	var product Product

	// Update - 更新多个字段 下面这句没有更新成功，要了老命了，对于空的零值的都没有设置成功
	//db.Model(&product).Updates(Product{Price: 200, Code: ""}) // 仅更新非零值字段 除了 Code 没有更新，其他非零值的字段都更新了
	//db.Model(&product).Updates(Product{Price: 0, Code: ""})   // 仅更新非零值字段
	// 允许更新零值
	db.Model(&product).Updates(Product{Price: 200, Code: sql.NullString{"", true}})
}

/**
UPDATE `products` SET `updated_at`='2021-08-21 22:57:23.357',`price`=200

UPDATE `products` SET `updated_at`='2021-08-21 22:57:23.357'

解决方法
在馍丁定义的时候指定 sql.NullString

为什么不允许设置呢？
因为go语言中string类型默认为空字符串，int类型默认是0，如果允许设置，明明只更新了一个price字段，就有可能将其他字段都设置成了零值，就真的要了老命了

UPDATE `products` SET `updated_at`='2021-08-21 23:07-form-validate:06-json-protobuf.013',`code`='',`price`=200

阻止全局更新
如果在没有任何条件的情况下执行批量更新，默认情况下，GORM 不会执行该操作，并返回 ErrMissingWhereClause 错误

对此，你必须加一些条件，或者使用原生 SQL，或者启用 AllowGlobalUpdate 模式，例如：
*/
