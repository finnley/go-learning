package main

import (
	"fmt"
	"os"
)

func main() {
	// os.Args 提供原始命令行参数访问功能。 注意，切片中的第一个参数是该程序的路径， 而 os.Args[1:]保存了程序全部的参数。
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}

/**
[C:\www\go-code\go-by-example\65-command-line-arguments\main.exe a b c d]
[a b c d]
c


要实验命令行参数，最好先使用 go build 编译一个可执行二进制文件
 */