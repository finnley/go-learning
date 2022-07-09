package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"net"
	"time"
)

func main() {
	//// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	//dsn := "root:root@tcp(127.0.0.1:3357)/gorm_test?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	dsn := "tcp(127.0.0.1:3357)/"
	db, err := sql.Open("mysql", dsn)
	fmt.Println(err)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("success")

	//d := db.Driver()
	//d.Open(dsn)

	//fmt.Println(tcpGather("127.0.0.1", []string{"3357", "3358"}))
	//cnf := mysql.NewConfig()
	//cnf.Addr = "127.0.0.1:3357"
	//conn, err := mysql.NewConnector(cnf)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//c, err := conn.Connect(context.Background())
	//if err != nil {
	//	fmt.Println("22:", err)
	//}
	//fmt.Println(c)
	cc, err := net.Dial("tcp", "111.231.87.78:3357")
	if err != nil {
		fmt.Println("22:", err)
	}
	fmt.Println(cc.RemoteAddr())
	fmt.Println(cc)

	//service, err := net.LookUpService(23)
	//netdb.ServiceByName("domain", "udp")
	//fmt.Println(netdb.GetServByPort(3357, &netdb.Protoent{}))
}

func testing() {
	//func (c *connector) Connect(ctx context.Context) (driver.Conn, error) {

}

func tcpGather(ip string, ports []string) map[string]string {
	// 检查 emqx 1883, 8083, 8080, 18083 端口

	results := make(map[string]string)
	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		// 3 秒超时
		conn, err := net.DialTimeout("tcp", address, 3*time.Second)
		fmt.Println(conn)
		if err != nil {
			results[port] = "failed"
			// todo log handler
		} else {
			if conn != nil {
				results[port] = "success"
				_ = conn.Close()
			} else {
				results[port] = "failed"
			}
		}
	}
	return results
}
