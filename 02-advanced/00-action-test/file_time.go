package main

import (
	"fmt"
	"github.com/djherbis/times"
	"log"
)

func main() {
	t, err := times.Stat("./redis_slow_log/slow.log")
	if err != nil {
		log.Fatal(err.Error())
		fmt.Println(err)
	} else {
		fmt.Println(11)
		fmt.Println(t.BirthTime())
	}
}
