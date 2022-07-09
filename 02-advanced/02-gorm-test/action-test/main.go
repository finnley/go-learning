package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID   uint
	Name string
	// 也是为了解决空字符串的问题
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
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

	var total int64
	// 获取全部记录
	var users []User
	//result := db.Find(&users)
	result := db.Where("name = ?", "jinzhu3").Find(&users)
	fmt.Println(result.RowsAffected) // 返回找到的记录数，相当于 `len(users)`
	db.Where("name = ?", "jinzhu3").Count(&total)

	var ids []uint
	for _, user := range users {
		//fmt.Println(user.ID)
		ids = append(ids, user.ID)
	}
	fmt.Println(ids)
	fmt.Println(total)

	//fmt.Println(result.RowsAffected) // 返回找到的记录数，相当于 `len(users)`
	//fmt.Println(result.Error)        // returns error
}

/**
// 通过 first 查询单个数据  Get the first record ordered by primary key
[rows:1] SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1

[rows:1] SELECT * FROM `users` WHERE `users`.`id` = 3 ORDER BY `users`.`id` LIMIT 1
3

*/
