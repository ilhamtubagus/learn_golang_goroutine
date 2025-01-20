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

func TestDisplayNumber(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}
	fmt.Println("Testing DisplayNumber...")

	time.Sleep(5 * time.Second)
}
