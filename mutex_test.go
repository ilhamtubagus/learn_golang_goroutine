package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestRaceConditionWithMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				mutex.Lock()
				x += 1
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(3 * time.Second)
	// Final value of x is 100,000. Race conditions are avoided using a mutex.
	fmt.Println("Final value of x:", x)
}
