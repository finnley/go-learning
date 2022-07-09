package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	fmt.Println(<-time.After(time.Second * 3))
	fmt.Println(time.Now())
	//time.Tick()
}
