package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)

	// Perform asynchronous task here
	time.Sleep(1 * time.Second)
	fmt.Println("Asynchronous task completed")
}

func TestAsynchronousTasks(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("All asynchronous tasks completed")
}
