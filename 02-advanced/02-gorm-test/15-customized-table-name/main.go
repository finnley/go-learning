package main

import (
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Language struct {
	gorm.Model
	Name    string
	AddTime time.Time // 每个记录创建的时候自动加上当前时间加入到 AddTime 中
	//AddTime sql.NullTime
}

// gorm中可以通过给某一个struct 添加 TableName 方法来自定义表名
//func (Language) TableName() string {
//	return "my_language"
//}

func (l *Language) BeforeCreate(tx *gorm.DB) (err error) {
	l.AddTime = time.Now()
	return
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
		// 2.NamingStrategy 这个配置不能和 Table 同时配置，如果同时配置会以 TableName 为准
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "gorm_",
		},
		Logger:            newLogger,
		AllowGlobalUpdate: true,
	})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Language{})

	db.Create(&Language{
		Name: "python",
	})
}

/**
[rows:1] INSERT INTO `gorm_languages` (`created_at`,`updated_at`,`deleted_at`,`name`,`add_time`) VALUES ('2021-08-22 23:41:02-init-router-ping-pong.244','2021-08-22 23:41:02-init-router-ping-pong.244',NULL,'python','2021-08-22 23:41:02-init-router-ping-pong.243')

*/
