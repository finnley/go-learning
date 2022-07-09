package main

import (
	"fmt"
	"time"
)

func chars(prefix string) {
	for c := 'A'; c <= 'Z'; c++ {
		fmt.Printf("%s: %c\n", prefix, c)
		time.Sleep(time.Microsecond)
	}
}

func main() {
	// goroutine
	go chars("goroutine")
	chars("main")

	time.Sleep(time.Second * 3)
}
