package main

import "fmt"

func Add(a, b int) int {
	total := a + b
	return total
}

func main() {
	total := Add(1, 2)
	fmt.Println(total)
}
