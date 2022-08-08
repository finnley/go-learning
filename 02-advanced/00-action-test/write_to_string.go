package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func write() {
	//创建一个新文件，写入内容 5 句 “http://c.biancheng.net/golang/”
	filePath := "./redis_slow_log/a.log"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for {
		write.WriteString("http://c.biancheng.net/golang/ \n")
		time.Sleep(1 * time.Second)
		write.Flush()
	}
	//Flush将缓存的文件真正写入到文件中

}

func main() {
	go write()

	os.Rename("redis_slow_log/a.log", "redis_slow_log/b.log")

	time.Sleep(10 * time.Second)
}
