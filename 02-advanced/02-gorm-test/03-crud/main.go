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

	// 定义一个表结构，将表结构直接生成对应的表
	// 迁移 schema
	_ = db.AutoMigrate(&Product{})

	// Create
	//db.Create(&Product{Code: "D42", Price: 100})
	//db.Create(&Product{Code: "D42", Price: 300})

	// Read
	var product Product
	//db.First(&product, 1)                 // 根据整形主键查找
	//db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(&product).Update("Price", 1000) // 更新的条件是上面查询出来的记录id作为条件
	// Update - 更新多个字段
	//db.Model(&product).Updates(Product{Price: 200, Code: "AAA"}) // 仅更新非零值字段 更新的条件是上面查询出来的记录id作为条件
	//db.Model(&product).Updates(Product{Price: 0, Code: "AAA"}) // 仅更新非零值字段 更新的条件是上面查询出来的记录id作为条件
	//db.Model(&product).Updates(Product{Price: 2000, Code: ""}) // 仅更新非零值字段 更新的条件是上面查询出来的记录id作为条件
	//db.Model(&product).Updates(Product{Price: 2000}) // 仅更新非零值字段 更新的条件是上面查询出来的记录id作为条件
	//db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "G42"}) // 更新的条件是上面查询出来的记录id作为条件
	// 下面map格式可以更新零值字段
	//db.Model(&product).Updates(map[string]interface{}{"Price": 0, "Code": ""}) // 更新的条件是上面查询出来的记录id作为条件

	//// Delete - 删除 product 逻辑删除
	//db.Delete(&product, 1)
}

/**
[1.748ms] [rows:1] INSERT INTO `products` (`created_at`,`updated_at`,`deleted_at`,`code`,`price`) VALUES ('2021-08-21 22:44:56.551','2021-08-21 22:44:56.551',NULL,'D42',100)

[0.844ms] [rows:1] SELECT * FROM `products` WHERE `products`.`id` = 1 AND `products`.`deleted_at` IS NULL ORDER BY `products`.`id` LIMIT 1

[0.671ms] [rows:1] SELECT * FROM `products` WHERE code = 'D42' AND `products`.`deleted_at` IS NULL AND `products`.`id` = 1 ORDER BY `products`.`id` LIMIT 1

[1.510ms] [rows:1] UPDATE `products` SET `price`=200,`updated_at`='2021-08-21 22:44:56.555' WHERE `id` = 1

[1.053ms] [rows:1] UPDATE `products` SET `updated_at`='2021-08-21 22:44:56.556',`code`='F42',`price`=200 WHERE `id` = 1

[0.852ms] [rows:1] UPDATE `products` SET `code`='F42',`price`=200,`updated_at`='2021-08-21 22:44:56.557' WHERE `id` = 1

[2.607ms] [rows:1] UPDATE `products` SET `deleted_at`='2021-08-21 22:44:56.558' WHERE `products`.`id` = 1 AND `products`.`id` = 1 AND `products`.`deleted_at` IS NULL
*/
