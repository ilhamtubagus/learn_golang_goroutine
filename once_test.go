package main

import (
	"fmt"
	"sync"
	"testing"
)

var counter int = 0

func OnlyOnce() {
	counter++
}

func TestOnlyOnce(t *testing.T) {
	once := sync.Once{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go func() {
			group.Add(1)

			once.Do(OnlyOnce)
			//OnlyOnce()

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter) // Expected: 1 because OnlyOnce invoked with once.Do
}
