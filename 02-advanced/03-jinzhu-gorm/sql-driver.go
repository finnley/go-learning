package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"

	//_ "github.com/go-sql-driver/mysql"
	"net"
)

type Config struct {
	IP   string
	Port int
}

type buffer struct {
	buf     []byte // buf is a byte buffer who's length and capacity are equal.
	nc      net.Conn
	idx     int
	length  int
	timeout time.Duration
}

type mysqlConn struct {
	buf              buffer
	netConn          net.Conn
	cfg              *mysql.Config
	maxAllowedPacket int
	maxWriteSize     int
	writeTimeout     time.Duration
	sequence         uint8
	parseTime        bool
}

type ServiceDiscoveryMySQLServiceDriver struct{}

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

func (d *ServiceDiscoveryMySQLServiceDriver) Open(dsn string) (driver.Conn, error) {
	cfg, err := mysql.ParseDSN(dsn)
	if err != nil {
		return nil, err
	}
	c := &connector{
		cfg: cfg,
	}
	_, err = c.Connect(context.Background())
	//buf := make([]byte, 0, 4096) // big buffer
	//tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	//for {
	//	n, err := conn.Read(tmp)
	//	if err != nil {
	//		if err != io.EOF {
	//			fmt.Println("read error:", err)
	//		}
	//		break
	//	}
	//	//fmt.Println("got", n, "bytes.")
	//	buf = append(buf, tmp[:n]...)
	//
	//}
	//fmt.Println("total size:", len(buf))
	//fmt.Println(string(buf))
	//if strings.Contains(string(buf), "mysql") {
	//	fmt.Println("mysql service discovery")
	//} else {
	//	fmt.Println("no service discovery")
	//}
	return nil, err
}

func main() {
	dsn := "root:root@tcp(127.0.0.1:3357)/gorm_test"
	//dsn := "@tcp(127.0.0.1:3357)/"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close() // 注意这行代码要写在上面err判断的下面

	err = db.Ping()

	//s := ServiceDiscoveryMySQLServiceDriver{}
	//_, err := s.Open(dsn)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println("success")
}
