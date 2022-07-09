package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3357)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"

	// 设置全局的 logger，这个 logger 在执行每个 sql 的时候会打印每一行 sql
	// 参考: https://gorm.io/zh_CN/docs/logger.html
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
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// 定义一个表结构，将表结构直接生成对应的表 migrations
	// 迁移 schema
	/**
	CREATE TABLE `products` (
		`id` BIGINT UNSIGNED AUTO_INCREMENT,
		`created_at` datetime ( 3 ) NULL,
		`updated_at` datetime ( 3 ) NULL,
		`deleted_at` datetime ( 3 ) NULL,
		`code` LONGTEXT,
		`price` BIGINT UNSIGNED,
	PRIMARY KEY ( `id` ),
	INDEX idx_products_deleted_at ( `deleted_at` ))
	*/
	_ = db.AutoMigrate(&Product{})
}

/**
[0.169ms] [rows:-] SELECT DATABASE()

[0.670ms] [rows:-] SELECT count(*) FROM information_schema.tables WHERE table_schema = 'gorm_test' AND table_name = 'products' AND table_type = 'BASE TABLE'

[3.623ms] [rows:0] CREATE TABLE `products` (`id` bigint unsigned AUTO_INCREMENT,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,`deleted_at` datetime(3) NULL,`code` longtext,`price` bigint unsigned,PRIMARY KEY (`id`),INDEX idx_products_deleted_at (`deleted_at`))

*/
