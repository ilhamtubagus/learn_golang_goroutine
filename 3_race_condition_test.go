package main

import (
	"fmt"
	"testing"
	"time"
)

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x += 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	// Final value of x should be 100,000, but it could be less due to race conditions.
	fmt.Println("Final value of x:", x)
}
