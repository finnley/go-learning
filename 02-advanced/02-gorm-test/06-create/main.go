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

type User struct {
	ID   uint
	Name string
	// Email 指针类型也是为了解决空字符串的问题
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

	// 定义一个表结构，将表结构直接生成对应的表
	// 迁移 schema
	/**
	CREATE TABLE `users` (
		`id` BIGINT UNSIGNED AUTO_INCREMENT,
		`name` LONGTEXT,
		`email` LONGTEXT,
		`age` TINYINT UNSIGNED,
		`birthday` datetime ( 3 ) NULL,
		`member_number` LONGTEXT,
		`activated_at` datetime ( 3 ) NULL,
		`created_at` datetime ( 3 ) NULL,
	`updated_at` datetime ( 3 ) NULL,
	PRIMARY KEY ( `id` ));
	*/
	_ = db.AutoMigrate(&User{})
	//db.Create(&User{
	//	Name: "江苏东台",
	//})

	//db.Model(&User{ID: 1}).Update("Name", "") // 可以更新零值
	//db.Model(&User{ID: 1}).Updates(User{Name: ""}) // 可更新零值

	//db.Create(&User{Name: "golang"}) // 会 Insert 所有字段，没有指定的字段会设置为默认值
	db.Create(&User{Name: ""})

	// 会更新零值字段 UPDATE `users` SET `name`='',`updated_at`='2021-09-26 07-form-validate:56:33.71' WHERE `id` = 1
	//db.Model(&User{ID: 1}).Update("Name", "")
	// Updates 语句不会更新零值，但是 Update 会更新
	//db.Model(&User{ID: 1}).Updates(User{Name: ""})
	//empty := ""
	//db.Model(&User{ID: 1}).Updates(User{Email: &empty})
	// 总结：
	//解决近更新非零值字段的方法有两种：
	//1. 将 string 设置为 *string
	//2. 使用 sql.NULLxxx

	//	user := User{
	//		Name: "php",
	//	}
	//	fmt.Println(user.ID)             // 0
	//	result := db.Create(&user)       // 通过数据的指针来创建
	//	fmt.Println(user.ID)             // 2 返回插入数据的主键
	//	fmt.Println(result.Error)        // <nil> 返回 error nil表示没有错误
	//	fmt.Println(result.RowsAffected) // 1 返回插入记录的条数
}

/**
CREATE TABLE `users` (`id` bigint unsigned AUTO_INCREMENT,`name` longtext,`email` longtext,`age` tinyint unsigned,`birthday` datetime(3) NULL,`member_number` longtext,`activated_at` datetime(3) NULL,`created_at` datetime(3) NULL,`updated_at` datetime(3) NULL,PRIMARY KEY (`id`))

db.Create(&User{Name: "golang"})
INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('golang',NULL,0,NULL,NULL,NULL,'2021-08-22 02-init-router-ping-pong:03-router-group:02-init-router-ping-pong.905','2021-08-22 02-init-router-ping-pong:03-router-group:02-init-router-ping-pong.905')

db.Model(&User{ID: 1}).Update("Name", "")
[rows:1] UPDATE `users` SET `name`='',`updated_at`='2021-08-22 02-init-router-ping-pong:17:19.699' WHERE `id` = 1

下面的语句并没有设置零值 Updates 语句不会更新零值，但是 Update 会更新
db.Model(&User{ID: 1}).Updates(User{Name: ""})
[rows:1] UPDATE `users` SET `updated_at`='2021-08-22 02-init-router-ping-pong:23:17.165' WHERE `id` = 1

// 下面方式更新是会设置成空字符串的
empty := ""
db.Model(&User{ID: 1}).Updates(User{Email: &empty})
[rows:1] UPDATE `users` SET `email`='',`updated_at`='2021-08-22 02-init-router-ping-pong:26:10.593' WHERE `id` = 1

// 总结：
解决近更新非零值字段的方法有两种：
1. 将 string 设置为 *string
2. 使用 sql.NULLxxx

0

[1.347ms] [rows:1] INSERT INTO `users` (`name`,`email`,`age`,`birthday`,`member_number`,`activated_at`,`created_at`,`updated_at`) VALUES ('php',NULL,0,NULL,NULL,NULL,'2021-08-22 02-init-router-ping-pong:32:39.668','2021-08-22 02-init-router-ping-pong:32:39.668')
3
<nil>
1

*/
