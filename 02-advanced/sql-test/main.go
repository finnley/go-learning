package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // 相当于在 import 的时候执行mysql驱动里面的 init()
	"time"
)

var db *sql.DB

func initMySQL() (err error) {
	// DSN:Data Source Name
	dsn := "root:root@tcp(127.0.0.1:3357)/gin_mysql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// 尝试与数据库简历连接（校验dsn是否正确）
	err = db.Ping()
	if err != nil {
		return err
	}
	db.SetConnMaxLifetime(time.Second * 10) // 最大连接存活时间
	db.SetMaxOpenConns(200)                 // 最大连接数
	db.SetMaxIdleConns(10)                  // 最大空闲连接数
	//orm_test.SetMaxOpenConns(2) // 最大连接数
	//orm_test.SetMaxIdleConns(1) // 最大空闲连接数
	//return nil
	return
}

type user struct {
	id   int
	name string
	age  int
}

// 查询单条数据实例
func queryRowDemo() {
	sqlStr := "select id, name, age from user where id=?"
	var u user
	// 非常重要：确保QueryRow之后调用 Scan 方法，否则持有的数据库连接不会被释放
	//err := orm_test.QueryRow(sqlStr, 1).Scan(&u.id, &u.name, &u.age)
	row := db.QueryRow(sqlStr, 1)
	//row = orm_test.QueryRow(sqlStr, 2)
	err := row.Scan(&u.id, &u.name, &u.age)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
}

// 查询多条数据
func queryMultiRowDemo() {
	sqlStr := "select id, name, age from user where id>?"
	rows, err := db.Query(sqlStr, 0)
	if err != nil {
		fmt.Printf("scan failed, err: %v\n", err)
		return
	}
	// 非常重要：关闭rows释放持有的数据链接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err: %v\n", err)
			return
		}
		fmt.Printf("id: %d, name: %s, age: %d\n", u.id, u.name, u.age)
	}
}

// 插入数据
func insertRowDemo() {
	sqlStr := "insert into user(name, age) values(?, ?)"
	ret, err := db.Exec(sqlStr, "peter", 10)
	if err != nil {
		fmt.Printf("insert failed, err: %v\n", err)
		return
	}
	theId, err := ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastinsert Id failed, err: %v\n", err)
		return
	}
	fmt.Printf("insert success, the id is %d\n", theId)
}

// 更新数据
func updateRowDemo() {
	sqlStr := "update user set name=? where id = ?"
	ret, err := db.Exec(sqlStr, "update", 4)
	if err != nil {
		fmt.Printf("update failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get lastinsert Id failed, err: %v\n", err)
		return
	}
	fmt.Printf("update success, affected rows: %d\n", n)
}

// 删除数据
func deleteRowDemo() {
	sqlStr := "delete from user where id = ?"
	ret, err := db.Exec(sqlStr, 4)
	if err != nil {
		fmt.Printf("delete failed, err: %v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get RowsAffected failed, err: %v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows: %d\n", n)
}

func main() {
	//// DSN:Data Source Name
	//dsn := "root:root@tcp(127.0.0.1:3357)/gin_mysql_test"
	//orm_test, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	panic(err)
	//}
	//// 做完错误检查之后，确保db不为nil
	//// Close 可以理解为释放掉数据库连接相关资源
	//defer orm_test.Close() // 注意这行diamante要卸载上面 err 判断的下面
	//
	//// 尝试与数据库简历连接（校验dsn是否正确）
	//err = orm_test.Ping()
	//if err != nil {
	//	fmt.Printf("connect to orm_test failed, err: %v\n", err)
	//	return
	//}
	//fmt.Println("connect to orm_test success")
	// 上面代码优化 使用init
	if err := initMySQL(); err != nil {
		fmt.Printf("connect to orm_test failed, err: %v\n", err)
		return
	}
	// 做完错误检查之后，确保db不为nil
	// Close 可以理解为释放掉数据库连接相关资源
	defer db.Close() // 注意这行diamante要卸载上面 err 判断的下面
	fmt.Println("connect to orm_test success")

	queryRowDemo()
	insertRowDemo()
	//updateRowDemo()
	deleteRowDemo()
	queryMultiRowDemo()
	fmt.Println("查询结束...")
}
