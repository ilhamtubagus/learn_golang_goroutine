package main

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello, World!")
}

func DisplayNumber(num int) {
	fmt.Println("Displaying number:", num)
}

func TestHelloWorld(t *testing.T) {
	go RunHelloWorld()
	fmt.Println("Testing HelloWorld...")

	time.Sleep(1 * time.Second)
}

// TestDisplayNumber tests the concurrent execution of DisplayNumber function.
// It launches 100,000 goroutines, each calling DisplayNumber with a unique integer.
// The test then waits for 5 seconds to allow goroutines to complete.
//
// Parameters:
//   - t: *testing.T - The testing object provided by the Go testing framework.
//
// This function does not return any value.
func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	fmt.Println("Testing DisplayNumber...")

	time.Sleep(5 * time.Second)
}
