package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// sync.Pool can be used to store temporary value that can be used in multiple goroutines
// Get() will retrieve the value and remove it from the pool, when Get() invoked later, it will create the default value from New()

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			// Will be invoked when the pool is empty and a new object is needed.
			return "Default"
		},
	}
	group := &sync.WaitGroup{}

	pool.Put("Ilham")
	pool.Put("Tubagus")
	pool.Put("Arfian")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			defer group.Done()
			data := pool.Get()
			fmt.Println("Data:", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	group.Wait()
	fmt.Println("Done") // Expected: 3
}
