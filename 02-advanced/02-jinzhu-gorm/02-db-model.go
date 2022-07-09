package main

import (
	"database/sql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"time"
)

type UserModel struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64 `gorm:"column:user_age"`
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        // 设置字段大小为255
	MemberNumber *string `gorm:"unique;not null"` // 设置会员号（member number）唯一并且不为空
	//Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 num 为自增类型
	Num      int    `gorm:"column:num"` // 设置 num 为自增类型
	Address  string `gorm:"index:addr"` // 给address字段创建名为addr的索引
	IgnoreMe int    `gorm:"-"`          // 忽略本字段
}

type TestUser struct {
	Id int
}

// 使用`AnimalID`作为主键
type Animal struct {
	AnimalID int64 `gorm:"primary_key"`
	Name     string
	Age      int64
}

func (Animal) TableName() string {
	return "animal"
}

func main() {
	// MySQL
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	//defer db.Close()

	// SQLite3
	db, err := gorm.Open("sqlite3", "/Users/finnley/Coding/go-learning/02-advanced/03-jinzhu-gorm/gorm.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SingularTable(true)

	// 创建表，自动迁移（把结构体和数据表进行对应）
	db.AutoMigrate(&UserModel{})
	db.AutoMigrate(&Animal{})

	// 创建表
	db.Table("test_user").CreateTable(&TestUser{})
}
