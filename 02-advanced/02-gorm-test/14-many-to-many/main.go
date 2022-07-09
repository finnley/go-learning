package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Student 拥有并属于多种 language，`student_languages` 是连接表
type Student struct {
	gorm.Model
	Languages []Language `gorm:"many2many:student_languages;"`
}

type Language struct {
	gorm.Model
	Name string
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

	db.AutoMigrate(&Student{})

	// 添加数据
	//languages := []Language{}
	//languages = append(languages, Language{Name: "go"})
	//languages = append(languages, Language{Name: "java"})
	//student := Student{
	//	Languages: languages,
	//}
	//db.Create(&student)

	// 获取数据
	var student Student
	db.Preload("Languages").First(&student)
	for _, language := range student.Languages {
		fmt.Println(language.Name)
	}
}

/**
[1.122ms] [rows:2] INSERT INTO `languages` (`created_at`,`updated_at`,`deleted_at`,`name`) VALUES ('2021-08-22 22:45:44.641','2021-08-22 22:45:44.641',NULL,'go'),('2021-08-22 22:45:44.641','2021-08-22 22:45:44.641',NULL,'java') ON DUPLICATE KEY UPDATE `id`=`id`

[1.390ms] [rows:2] INSERT INTO `user_languages` (`user_id`,`language_id`) VALUES (1,1),(1,2) ON DUPLICATE KEY UPDATE `user_id`=`user_id`

[5.738ms] [rows:1] INSERT INTO `users` (`created_at`,`updated_at`,`deleted_at`) VALUES ('2021-08-22 22:45:44.64','2021-08-22 22:45:44.64',NULL)


*/
