package main

import (
	"flag"
	"fmt"
)

var (
	intflag    int
	boolflag   bool
	stringflag string
)

func init() {
	flag.IntVar(&intflag, "intflag", 0, "int flag value")
	flag.BoolVar(&boolflag, "boolflag", false, "bool flag value")
	flag.StringVar(&stringflag, "stringflag", "default", "string flag value")
}

func main() {
	flag.Parse()

	fmt.Println("init flag: ", intflag)
	fmt.Println("bool flag: ", boolflag)
	fmt.Println("string flag: ", stringflag)
}

/**
 ✗ go build -o main main.go
 ✗ ./main -intflag 12 -boolflag 1 -stringflag test
init flag:  12
bool flag:  true
string flag:

// 默认值：
✗ ./main -intflag 12 -boolflag 1
init flag:  12
bool flag:  true
string flag: default

*/
