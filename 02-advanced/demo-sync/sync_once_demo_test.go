package demo_sync

import (
	"fmt"
	"sync"
	"testing"
)

func TestSyncOnce(t *testing.T) {
	var once sync.Once
	onceFunc := func() {
		fmt.Println("Only once")
	}

	for i := 0; i < 10; i++ {
		once.Do(onceFunc)
	}
}
