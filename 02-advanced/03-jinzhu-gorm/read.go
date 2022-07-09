package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
	"strings"
)

func main2() {
	//conn, err := net.Dial("tcp", "google.com:80")
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")

	buf := make([]byte, 0, 4096) // big buffer
	tmp := make([]byte, 256)     // using small tmo buffer for demonstrating
	for {
		fmt.Println(333)
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
			}
			break
		}
		//fmt.Println("got", n, "bytes.")
		buf = append(buf, tmp[:n]...)

	}
	fmt.Println(111111)
	fmt.Println("total size:", len(buf))
	fmt.Println(string(buf))
	if strings.Contains(string(buf), "mysql") {
		fmt.Println("mysql service discovery")
	} else {
		fmt.Println("no service discovery")
	}
}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("dial error:", err)
		return
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	var buf bytes.Buffer
	io.Copy(&buf, conn)
	fmt.Println("total size:", buf.Len())
}
