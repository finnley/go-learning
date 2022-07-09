package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Golang里面数据路创建之后底层都是连接池，而不是单个连接
	dsn := "root:root@tcp(192.168.1.8:3357)/gin_mysql_test?charset=utf8mb4&loc=PRC&parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
	}
	// ping是从连接池里面获取一个连接，然后检查下状态
	fmt.Println(db.Ping())

	// 执行
	fmt.Println(db.Exec(`
create table if not exists test1(
    id bigint primary key auto_increment,
    name varchar(32) not null default '' comment '名称'
) engine=innodb default charset utf8mb4;
`))

	//// 更新
	//	result, err := orm_test.Exec(`
	//update test1 set name="finnley222" where id = 2
	//`)
	//	fmt.Println(result.LastInsertId())
	//	fmt.Println(result.RowsAffected())

	//// 插入
	//	resInsert, err := orm_test.Exec(`
	//insert test1(name) values("gerry");
	//`)
	//	fmt.Println(resInsert.LastInsertId())
	//	fmt.Println(resInsert.RowsAffected())

	// 查询
	rows, err := db.Query("select id, name from test1")
	var (
		id   int
		name string
	)
	for rows.Next() {
		rows.Scan(&id, &name)
		fmt.Println(id, name)
	}

}
