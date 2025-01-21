package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGomaxprocs(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Printf("Total CPU cores: %d\n", totalCpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Printf("Total Thread: %d\n", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Printf("Total Goroutine: %d\n", totalGoroutine)

	wg.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			wg.Done()
		}()
	}

	totalCpu := runtime.NumCPU()
	fmt.Printf("Total CPU cores: %d\n", totalCpu)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Printf("Total Thread: %d\n", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Printf("Total Goroutine: %d\n", totalGoroutine)

	wg.Wait()
}
