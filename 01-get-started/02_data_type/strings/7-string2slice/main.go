package main

import "fmt"

func main() {
	s := "abc"
	b := []byte(s)
	s2 := string(b)

	fmt.Println(b)  // [97 98 99]
	fmt.Println(s2) // abc
}
