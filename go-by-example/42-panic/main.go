package main

import "os"

func main() {
	panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}

/**
panic: a problem

goroutine 1 [running]:
main.main()
	C:/www/go-code/go-by-example/42-panic/fibonacci.go:6 +0x40

 */