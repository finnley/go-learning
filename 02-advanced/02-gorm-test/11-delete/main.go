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

type NewUser struct {
	ID     uint
	MyName string `gorm:"column:name"`
	// 也是为了解决空字符串的问题
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString
	ActivatedAt  sql.NullTime
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Deleted      gorm.DeletedAt
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

	db.AutoMigrate(&NewUser{})

	//var users = []NewUser{{MyName: "jinzhu1"}, {MyName: "jinzhu2"}, {MyName: "jinzhu3"}}
	//db.Create(&users)

	//for _, admin := range users {
	//	fmt.Println(admin.ID)
	//}

	db.Delete(&NewUser{}, 1)

	var users []NewUser
	db.Find(&users)
	for _, user := range users {
		fmt.Println(user.ID)
	}

	// 硬删除
	db.Unscoped().Delete(&NewUser{ID: 2})
}

/**
db.Delete(&NewUser{}, 1)
[rows:1] UPDATE `new_users` SET `deleted`='2021-08-22 14:43:24.325' WHERE `new_users`.`id` = 1 AND `new_users`.`deleted` IS NULL

// 硬删除
db.Unscoped().Delete(&NewUser{ID: 2})
[rows:1] DELETE FROM `new_users` WHERE `new_users`.`id` = 2
*/
