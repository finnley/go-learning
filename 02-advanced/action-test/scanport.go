package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	//host := "111.231.87.78"
	//host := "127.0.0.1"
	//port := 3380
	host := "10.186.60.76"
	port := 3310
	conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", host, port), 30*time.Second)
	//tcpAddr := net.TCPAddr{
	//	IP:   net.ParseIP(host),
	//	Port: port,
	//}
	//conn, err := net.DialTCP("tcp", nil, &tcpAddr)
	fmt.Println(err)

	if err == nil {
		fmt.Println("success")
		conn.Close()
		conn = nil
	} else {
		fmt.Println("failed")
	}
	hostAddr, err := net.LookupAddr(host)
	if err != nil {
		fmt.Printf("service_probe can not get hostname %v, err: %v", host, err)
	}
	fmt.Println(hostAddr)
}

func main4() {
	ips := "172.20.134.7/20"
	hostList, err := ParseIpsStr(ips)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(hostList)
}
