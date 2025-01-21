package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64 = 0
	wg := &sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		go func() {
			wg.Add(1)
			for j := 0; j < 100; j++ {
				atomic.AddInt64(&x, 1)
				// x+=1 -> will result in race conditions
			}
			wg.Done()
		}()
	}

	wg.Wait()
	// Final value of x should be 100,000, but it could be less due to race conditions.
	fmt.Println("Final value of x:", x)
}
