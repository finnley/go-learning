package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Worker %d starting\n", id)

	time.Sleep(time.Second)
	fmt.Printf("Workder %d done\n", id)
}

/**
Worker 1 starting
Worker 4 starting
Worker 5 starting
Worker 2 starting
Worker 3 starting
Workder 1 done
Workder 2 done
Workder 3 done
Workder 4 done
Workder 5 done

 */

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go worker(i, &wg)
	}

	wg.Wait()
}
