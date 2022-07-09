package main

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 参考: https://gorm.io/zh_CN/docs/models.html
type User struct {
	UserID uint `gorm:"primaryKey"`
	//Name   string `gorm:"column:user_name;type:varchar(50)"`
	//Name string `gorm:"column:user_name;type:varchar(50);index:idx_user_name;unique;default:'golang'"`
	//Name string `gorm:"column:user_name;type:varchar(50);index:idx_user_name;unique;default:''"`
	Name string `gorm:"column:user_name;type:varchar(50);index:idx_user_name;unique;not null:NOT NULL;default:''"`
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
	/**
	CREATE TABLE `users` (
		`user_id` BIGINT UNSIGNED AUTO_INCREMENT,
		`user_name` VARCHAR ( 50 ) UNIQUE DEFAULT '',
	PRIMARY KEY ( `user_id` ),
	INDEX idx_user_name ( `user_name` ));

	CREATE TABLE `users` (
		`user_id` BIGINT UNSIGNED AUTO_INCREMENT,
		`user_name` VARCHAR ( 50 ) NOT NULL UNIQUE DEFAULT '',
	PRIMARY KEY ( `user_id` ),
	INDEX idx_user_name ( `user_name` ));
	*/
	_ = db.AutoMigrate(&User{})

	db.Create(&User{})
}
