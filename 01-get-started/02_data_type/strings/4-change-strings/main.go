package main

import "fmt"

func main() {
	s := "left foot"
	t := s
	s += ", right foot"

	fmt.Println(s) // "left foot, right foot"
	fmt.Println(t) // "left foot"
}
