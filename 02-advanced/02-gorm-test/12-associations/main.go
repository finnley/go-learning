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

// `Employee` 属于 `Company`，`CompanyID` 是外键
type Employee struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company
}

type Company struct {
	ID   int
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

	db.AutoMigrate(&Employee{})

	//db.Create(&Employee{
	//	Name: "Tom",
	//	Company: Company{
	//		Name: "Tech",
	//	},
	//})

	//再新建
	//db.Create(&Employee{
	//	Name: "Jack",
	//	Company: Company{
	//		ID: 1,
	//	},
	//})

	// 关联查询-预加载
	var employee Employee
	db.First(&employee)
	fmt.Println(employee.Name)
	// 两行表没有做 join
	fmt.Println(employee.Name, employee.Company.Name, employee.Company.ID) // Tom "" 0 只取出了 employee 的信息，没有取出 company 信息

	// 1. 预加载
	db.Preload("Company").First(&employee)
	fmt.Println(employee.Name, employee.Company.Name, employee.Company.ID) // Tom Tech 1

	// 2. 关联查询-Joins 和上面 预加载功能一样
	db.Joins("Company").First(&employee)
	fmt.Println(employee.Name, employee.Company.Name, employee.Company.ID) // Tom Tech 1
}

/**
db.Preload("Company").First(&admin)
fmt.Println(admin.Name, admin.Company.ID)
[0.812ms] [rows:1] SELECT * FROM `companies` WHERE `companies`.`id` = 1

[1.901ms] [rows:1] SELECT * FROM `users` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 2 ORDER BY `users`.`id` LIMIT 1
golang 1

db.Joins("Company").First(&admin)
fmt.Println(admin.Name, admin.Company.ID)
[rows:1] SELECT `users`.`id`,`users`.`created_at`,`users`.`updated_at`,`users`.`deleted_at`,`users`.`name`,`users`.`company_id`,`Company`.`id` AS `Company__id`,`Company`.`name` AS `Company__name` FROM `users` LEFT JOIN `companies` `Company` ON `users`.`company_id` = `Company`.`id` WHERE `users`.`deleted_at` IS NULL AND `users`.`id` = 2 ORDER BY `users`.`id` LIMIT 1
golang 1

*/
