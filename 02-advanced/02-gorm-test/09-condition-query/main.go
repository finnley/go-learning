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
	Name string `gorm:"column:name"`
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

	// 通过 where 查询
	var admin User
	// 大小写不敏感
	db.Where("name = ?", "golang").First(&admin) // WHERE name = 'golang' ORDER BY `users`.`id` LIMIT 1
	db.Where("Name = ?", "golang").First(&admin) // WHERE Name = 'golang' AND `users`.`id` = 1 ORDER BY `users`.`id` LIMIT 1
	// 如果需要设置 大小写敏感，可以设置 column
	db.Where(&User{Name: "golang"}).First(&admin)

	var users []User
	//db.Where(&User{Name: "jinzhu1"}).Find(&users)
	// map 不会将 Name 转成 name，另外 map 可以查询 零值 字段
	db.Where(map[string]interface{}{"name": "jinzhu", "age": 18}).Find(&users)
	for _, user := range users {
		fmt.Println(user.ID)
	}
}

/**
[rows:1] SELECT * FROM `users` WHERE name = 'golang' ORDER BY `users`.`id` LIMIT 1


总结：
查询方式有3种
1. string 最灵活
2. struct 这种有坑，会屏蔽零值
3. map 可读性好

优先选择第二种，其次第三种，最后第一种
*/
