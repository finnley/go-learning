package main

import "fmt"

func main() {
	a := fmt.Errorf("current password does not meet password security policy")
	fmt.Println(a)
}
