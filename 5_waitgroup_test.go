package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup) {
	// Defer will be invoked at the end of the function regardless of whether the function returns normally or not.
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
