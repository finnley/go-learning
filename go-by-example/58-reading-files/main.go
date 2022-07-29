package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件需要经常进行错误检查， 这个帮助方法可以精简下面的错误检查过程
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 最基本的文件读取任务或许就是将文件内容读取到内存中
	dat, err := ioutil.ReadFile("./dat")
	check(err)
	// panic: open ./dat: no such file or directory
	fmt.Print(string(dat))

	// 您通常会希望对文件的读取方式和内容进行更多控制。 对于这个任务，首先使用 Open 打开一个文件，以获取一个 os.File 值
	f, err := os.Open("./dat")
	check(err)

	// 从文件的开始位置读取一些字节。 最多允许读取 5 个字节，但还要注意实际读取了多少个。
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Seek 到一个文件中已知的位置，并从这个位置开始读取。
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: ", n2, o2)
	fmt.Printf("%v\n", string(b2[:n2]))

	// 例如，io 包提供了一个更健壮的实现 ReadAtLeast，用于读取上面那种文件。
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内建的倒带，但是 Seek(0, 0) 实现了这一功能
	_, err = f.Seek(0, 0)
	check(err)

	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	// 任务结束后要关闭这个文件 （通常这个操作应该在 Open 操作后立即使用 defer 来完成）。
	f.Close()
}

/**
hello
go
5 bytes: hello
2 bytes @ 6: go
2 bytes @ 6: go
5 bytes: hello
*/
