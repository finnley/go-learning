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

func main2() {
	flag.Parse()

	fmt.Println("init flag: ", intflag)
	fmt.Println("bool flag: ", boolflag)
	fmt.Println("string flag: ", stringflag)
}

/**
 ✗ go build -o main flag_type_var.go
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

func main() {
	flag.Parse()

	fmt.Println(flag.Args())
	fmt.Println("Non-Flag Argument Count:", flag.NArg())
	for i := 0; i < flag.NArg(); i++ {
		fmt.Printf("Argument %d: %s\n", i, flag.Arg(i))
	}

	fmt.Println("Flag Count:", flag.NFlag())
}

/**
✗ ./main -intflag 12 -- -stringflag test
[-stringflag test]
Non-Flag Argument Count: 2
Argument 0: -stringflag
Argument 1: test
Flag Count: 1
*/
