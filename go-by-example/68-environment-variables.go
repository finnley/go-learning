package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// 用 os.Setenv 来设置一个键值对。 使用 os.Getenv获取一个键对应的值。 如果键不存在，将会返回一个空字符串。
	os.Setenv("FOO", "1")
	fmt.Println("FOO:", os.Getenv("FOO"))               // FOO: 1
	fmt.Println("BAR:", os.Getenv("BAR"))               // BAR:
	fmt.Println(os.Getenv("PATH"))                      // /usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/Users/finnley/go/bin
	fmt.Println(os.Getenv("PATH") + ":/sbin:/usr/sbin") // /usr/bin:/bin:/usr/sbin:/sbin:/usr/local/go/bin:/Users/finnley/go/bin:/sbin:/usr/sbin

	fmt.Println(os.Getenv("LC_ALL"))

	fmt.Println()
	// 使用 os.Environ 来列出所有环境变量键值对。 这个函数会返回一个 KEY=value 形式的字符串切片。 你可以使用 strings.SplitN 来得到键和值。这里我们打印所有的键。
	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)
		fmt.Println(pair[0])
	}
}
