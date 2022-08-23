package main

import (
	"flag"
	"fmt"
)

var (
	intflag2    *int
	boolflag2   *bool
	stringflag2 *string
)

func init() {
	intflag2 = flag.Int("intflag", 0, "int flag value")
	boolflag2 = flag.Bool("boolflag", false, "bool flag value")
	stringflag2 = flag.String("stringflag", "default", "string flag value")
}

func main() {
	flag.Parse()

	fmt.Println("int flag:", *intflag2)
	fmt.Println("bool flag:", *boolflag2)
	fmt.Println("string flag:", *stringflag2)
}
