package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 这里展示了如何写入一个字符串（或者只是一些字节）到一个文件
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("./dat1", d1, 0644)
	check(err)

	// 对于更细粒度的写入，先打开一个文件
	f, err := os.Create("./dat2")
	check(err)

	// 打开文件后，一个习惯性的操作是：立即使用 defer 调用文件的 Close。
	defer f.Close()

	// 可以按期望的那样 Write 字节切片
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n", n2) // wrote 5 bytes

	n3, err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n", n3) // wrote 7 bytes

	// 调用 Sync 将缓冲区的数据写入硬盘。
	f.Sync()

	// 我们前面看到的带缓冲的 Reader 一样，bufio 还提供了的带缓冲的 Writer。
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	fmt.Printf("wrote %d bytes\n", n4)

	// 使用 Flush 来确保，已将所有的缓冲操作应用于底层 writer。
	w.Flush()
}
