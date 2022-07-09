package main

import (
	"fmt"
	"net"
)

func main() {
	host, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		return
	}
	fmt.Println(host)
}
