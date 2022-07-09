package main

import (
	"context"
	"database/sql/driver"
	"fmt"
	"github.com/go-sql-driver/mysql"
	//_ "github.com/go-sql-driver/mysql"
	"net"
)

//type Config struct {
//	IP string
//	Port int
//}

type connector struct {
	cfg *mysql.Config // immutable private copy.
}

func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {
	_, err := net.Dial("tcp", "111.231.87.78:3357")
	if err != nil {
		fmt.Println("22:", err)
	}
	return nil, err
}

type ServiceDiscoveryMySQLServiceDriver struct{}

func (d *ServiceDiscoveryMySQLServiceDriver) Open(dsn string) (driver.Conn, error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	c := &connector{
		cfg: cfg,
	}
	_, err = c.Connect(context.Background())
	return nil, err
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3357)/gorm_test"
	//db, err := sql.Open("mysql", dsn)
	//if err != nil {
	//	panic(err)
	//}
	//defer db.Close() // 注意这行代码要写在上面err判断的下面

	//err = db.Ping()
	s := ServiceDiscoveryMySQLServiceDriver{}
	_, err := s.Open(dsn)
	if err != nil {
		panic(err)
	}

	//fmt.Println("success")
}
