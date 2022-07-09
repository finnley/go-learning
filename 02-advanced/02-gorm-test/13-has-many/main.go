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

//05-type Owner struct {
//	gorm.Model
//	CreditCards []CreditCard
//}
//
//05-type CreditCard struct {
//	gorm.Model
//	Number  string
//	OwnerId uint
//}

// 下面定义指明外键，这样设置在数据表中不会添加外键
type Owner struct {
	gorm.Model
	CreditCards []CreditCard `gorm:"foreignKey:UserRefer"`
}

type CreditCard struct {
	gorm.Model
	Number    string
	UserRefer uint
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

	db.AutoMigrate(&CreditCard{})
	db.AutoMigrate(&Owner{})

	owner := Owner{}
	db.Create(&owner)
	db.Create(&CreditCard{
		Number:    "12",
		UserRefer: owner.ID,
	})
	db.Create(&CreditCard{
		Number:    "34",
		UserRefer: owner.ID,
	})
	var owner2 Owner
	db.Preload("CreditCards").First(&owner2)
	for _, card := range owner2.CreditCards {
		fmt.Println(card.Number)
	}
}
