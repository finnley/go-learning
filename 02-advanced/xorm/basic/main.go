package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"xorm.io/xorm"
)

const DriverName = "mysql"
const MasterDataSourceName = "root:root@tcp(192.168.33.33:3357)/xorm_test?charset=utf8"

type UserInfo struct {
	Id         int `xorm:"not null pk autoincr"`
	Name       string
	SysCreated int
	SysUpdated int
}

var engine *xorm.Engine

func main() {
	engine = newEngin()

	//execute()
	//ormInsert()

	//query()
	//ormGet()
	ormGetCols()
	//ormFindRows()

}

func newEngin() *xorm.Engine {
	engine, err := xorm.NewEngine(DriverName, MasterDataSourceName)
	if err != nil {
		log.Fatal(newEngin, err)
		return nil
	}
	// Debug 模式，打印全部的SQL语句，帮助对比，看ORM与SQL执行的对照关系
	engine.ShowSQL(true)
	return engine
}

// 通过query方式查询
func query() {
	sql := "SELECT * FROM user_info"
	//results, err := engine.Query(sql)
	//results, err := engine.QueryInterface(sql)
	results, err := engine.QueryString(sql)
	if err != nil {
		log.Fatal("query", sql, err)
		return
	}
	total := len(results)
	if total == 0 {
		fmt.Println("没有任何数据", sql)
	} else {
		for i, data := range results {
			fmt.Printf("%d = %v\n", i, data)
		}
	}
}

// 根据models的结构读取数据
func ormGet() {
	UserInfo := &UserInfo{Id: 2}
	ok, err := engine.Get(UserInfo)
	if ok {
		fmt.Printf("%v\n", *UserInfo)
	} else if err != nil {
		log.Fatal("ormGet error", err)
	} else {
		fmt.Println("ormGet empty id=", UserInfo.Id)
	}
}

// 获取指定字段
func ormGetCols() {
	UserInfo := &UserInfo{Id: 2}
	ok, err := engine.Cols("name").Get(UserInfo)
	if ok {
		fmt.Printf("%v\n", UserInfo)
	} else if err != nil {
		log.Fatal("ormGetCols error", err)
	} else {
		fmt.Println("ormGetCols empty id=2", UserInfo.Id)
	}
}

// 通过 execute 方法执行更新
func execute() {
	sql := `INSERT INTO user_info values(NULL, 'name', 0, 0)`
	affected, err := engine.Exec(sql)
	if err != nil {
		log.Fatal("execute error", err)
	} else {
		id, _ := affected.LastInsertId()
		rows, _ := affected.RowsAffected()
		fmt.Println("execute id = ", id, ", rows = ", rows)
	}
}

// 根据 Models 的接口映射数据表
func ormInsert() {
	UserInfo := &UserInfo{
		Id:         0,
		Name:       "梅西",
		SysCreated: 0,
		SysUpdated: 0,
	}
	id, err := engine.Insert(UserInfo)
	if err != nil {
		log.Fatal("ormInsert error", err)
	} else {
		fmt.Println("ormInsert id = ", id)
		fmt.Printf("%v\n", *UserInfo)
	}
}

func ormFindRows() {
	list := make([]UserInfo, 0)
	//list := make(map[int]UserInfo)
	//err := engine.Find(&list)
	//err := engine.Where("id>?", 1).Limit(100, 0).Find(&list)
	err := engine.Cols("id", "name").Where("id>?", 0).Limit(10).Asc("id", "sys_created").Find(&list)

	//list := make([]map[string]string, 0)
	//err := engine.Table("star_info").Cols("id", "name_zh", "name_en").Where("id>?", 1).Find(&list)

	if err == nil {
		fmt.Printf("%v\n", list) // [{1 张三 0 0} {2 李四 0 0} {3 王五 0 0}]
	} else {
		log.Fatal("ormFindRows error", err)
	}
}
