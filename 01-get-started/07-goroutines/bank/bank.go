package main

import (
	"fmt"
	"time"
)

var balance int

func Deposit(amount int) {
	balance = balance + amount
}

func Balance() int {
	return balance
}

func main() {
	//fmt.Println(Balance())
	//Deposit(100)
	//fmt.Println(Balance())


	// Alice:
	go func() {
		Deposit(200) // A1
		fmt.Println("=", Balance()) // A2
	}()

	// Bob:
	go Deposit(100) // B

	time.Sleep(2 * time.Second)
}
