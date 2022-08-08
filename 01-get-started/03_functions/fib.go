// 函数是一等公民，可以给函数实现接口
package main

import "bufio"

import (
	"fmt"
	"io"
	"strings"
)

// 斐波那契数列
// 先定义斐波那契数列生成器
// 1, 1, 2, 3, 5, 8, 13, 21
//	  a, b
// a和b往后挪了以下
// 		 a, b
//func fibonacci() func() int {
//	a, b := 0, 1
//	return func() int {
//		a, b = b, a+b
//		return a
//	}
//}

func printFileContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func fibonacci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

type intGen func() int

// 函数也可以事先接口
func (g intGen) Read(p []byte) (n int, err error) {
	// 获取下一个元素
	next := g()

	if next > 10000 {
		return 0, io.EOF
	}

	// 把下一个元素写入到 p 这个 byte 里面去，这个写起来比较底层，所以就找办法找个代理，找个已经实现了reader的代理
	s := fmt.Sprintf("%d\n", next)

	// TODO: incorrect if p is too small!
	return strings.NewReader(s).Read(p)
}

func main() {
	f := fibonacci()

	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 1
	//fmt.Println(f()) // 2
	//fmt.Println(f()) // 3
	//fmt.Println(f()) // 5
	//fmt.Println(f()) // 8
	//fmt.Println(f()) // 13
	//fmt.Println(f()) // 21
	printFileContents(f)
}
