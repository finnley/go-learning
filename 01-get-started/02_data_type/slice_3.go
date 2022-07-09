package main

import "fmt"

func main() {
	varlen()
}

func varlen() {
	s2 := make([]int, 4, 6)
	s2[0] = 1
	s2[1] = 2
	s2[2] = 3
	s2[3] = 4

	fmt.Printf("%p\n", &s2)

	s2 = append(s2, s2...)
	fmt.Printf("%p\n", &s2)
	fmt.Println(s2)
}
