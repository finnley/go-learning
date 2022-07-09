package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	var locker sync.RWMutex
	// A 写
	go func() {
		defer wg.Done()
		log.Println("A: Lock Before")
		locker.Lock()

		log.Println("A: Locked")
		time.Sleep(time.Second * 5)

		locker.Unlock()
		log.Println("A: Unlock")
	}()
	// B 读
	go func() {
		defer wg.Done()
		log.Println("B: Lock Before")
		locker.RLock()

		log.Println("B: Locked")
		time.Sleep(time.Second * 5)

		locker.RUnlock()
		log.Println("B: Unlock")
	}()
	wg.Wait()
}

/**
写 读 RLock 和 Lock 互斥
写 写 Lock 和 Lock 互斥
读 读 RLock 和 RLock 不互斥
*/
