package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		token := scanner.Text()
		fmt.Println(token)
	}
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
