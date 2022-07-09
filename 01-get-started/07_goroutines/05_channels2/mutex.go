package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var locker sync.Mutex
	go func() {
		defer wg.Done()
		log.Println("A: Lock Before")
		locker.Lock()

		log.Println("A: Locked")
		time.Sleep(time.Second * 5)

		locker.Unlock()
		log.Println("A: Unlock")
	}()
	go func() {
		defer wg.Done()
		log.Println("B: Lock Before")
		locker.Lock()

		log.Println("B: Locked")
		time.Sleep(time.Second * 5)

		locker.Unlock()
		log.Println("B: Unlock")
	}()
	wg.Wait()
}

/**
2022/03/24 19:41:50 B: Lock Before
2022/03/24 19:41:50 B: Locked
2022/03/24 19:41:50 A: Lock Before
2022/03/24 19:41:55 B: Unlock
2022/03/24 19:41:55 A: Locked
2022/03/24 19:42:00 A: Unlock
*/
