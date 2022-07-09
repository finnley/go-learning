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

	_ = db.AutoMigrate(&User{})

	//var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	var users = []User{
		{Name: "\u003e"},
		{Name: "\\u003e"},
		{Name: `\u003e`},
		{Name: "redis:mem:usage / on(id) group_left redis:info_all:maxmemory * 100 \\u003e= 80"},
		{Name: `redis:mem:usage / on(id) group_left redis:info_all:maxmemory * 100 \u003e= 80`},
		{Name: `redis:mem:usage / on(id) group_left redis:info_all:maxmemory * 100 \\u003e= 80`},
	}
	db.Create(&users)
	for _, user := range users {
		fmt.Println(user.Name)
	}

	//// 数量为 100 即使批量插入3条记录，但是一次性提交insert的是2条，也就是每2条提交一次sql记录
	//db.CreateInBatches(&users, 2)

	// 注意： 根据 map 创建记录时，association 不会被调用，且主键也不会自动填充
	//db.Model(&User{}).Create(map[string]interface{}{
	//	"Name": "jinzhu", "Age": 18,
	//})
	//
	//for _, user := range users {
	//	fmt.Println(user.ID) // 1,2,3
	//}
}

/**
db.Create(&users)
[2.651ms] [rows:3] INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('jinzhu1',NULL,0,NULL,NULL,NULL,'2021-08-22 07-form-validate:54:59.974','2021-08-22 07-form-validate:54:59.974'),('jinzhu2',NULL,0,NULL,NULL,NULL,'2021-08-22 07-form-validate:54:59.974','2021-08-22 07-form-validate:54:59.974'),('jinzhu3',NULL,0,NULL,NULL,NULL,'2021-08-22 07-form-validate:54:59.974','2021-08-22 07-form-validate:54:59.974')
4
5
6

db.CreateInBatches(users, 1)
[rows:1] INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('jinzhu2',NULL,0,NULL,NULL,NULL,'2021-08-22 07-form-validate:57:38.452','2021-08-22 07-form-validate:57:38.452')

[1.059ms] [rows:1] INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('jinzhu3',NULL,0,NULL,NULL,NULL,'2021-08-22 07-form-validate:57:38.453','2021-08-22 07-form-validate:57:38.453')
7
8
9

CreateInBatches 为什么不一次性提交所有的，还要分批次，这个是因为sql语句有长度限制

[rows:1] INSERT INTO `users` (`age`,`name`) VALUES (18,'jinzhu')
*/
