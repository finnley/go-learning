package main

import (
	"fmt"
	"os"
	"path/filepath"
	"syscall"
)

func main2() {
	//file, err := os.Stat("./redis_slow_log/slow.log")
	file, err := os.Stat("./redis_slow_log/2.log.go")
	//file, err := os.Stat("./redis_slow_log/")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(file.Name())
	//fmt.Printf("%+v\n", file.Sys().(*syscall.Stat_t).Atimespec.Sec)
	//fmt.Printf("%+v\n", file.Sys().(*syscall.Stat_t).Ctimespec.Sec)
	//fmt.Printf("%+v\n", file.Sys().(*syscall.Stat_t).Mtimespec.Sec)
	//fmt.Printf("%+v\n", file.Sys().(*syscall.Stat_t).Birthtimespec.Sec)
	fmt.Printf("%+v\n", file.Sys().(*syscall.Stat_t).Birthtimespec)

	_path := filepath.Join(`./a`, `b/c`, `../d/`)
	println(_path)
	slogLogFilePath := filepath.Join("./redis_slow_log", "aaa", "slow.log")
	fmt.Println(slogLogFilePath)
}

func main23() {
	file, _ := os.Stat("./redis_slow_log/06.log")
	fmt.Printf("%+v\n", file.Sys().(*syscall.Stat_t).Birthtimespec)

	os.Rename("redis_slow_log/06.log", "redis_slow_log/07.log")

	file, _ = os.Stat("./redis_slow_log/07.log")
	fmt.Printf("%+v\n", file.Sys().(*syscall.Stat_t).Birthtimespec)
}

func main() {
	s := make([]string, 4)
	fmt.Println(len(s), cap(s))
	s = append(s, "555")
	fmt.Printf("%+v, %d, %d\n", s, len(s), cap(s))
}
