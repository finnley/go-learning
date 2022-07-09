package main

import (
	"fmt"
	"net"
)

func main() {
	//host := "111.231.87.78"
	//host := "127.0.0.1"
	host := "10.186.62.66"
	//host := "50.62.227.1"
	hostAddr, err := net.LookupAddr(host)
	fmt.Println(hostAddr, err)
}
