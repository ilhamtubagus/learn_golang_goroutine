package main

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, value int, wg *sync.WaitGroup) {
	defer wg.Done()

	wg.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go AddToMap(data, i, group)
	}

	group.Wait()

	data.Range(func(key, value interface{}) bool {
		fmt.Println(key, ":", value)

		// Return true will continue iterating over the map,
		return true
	})
}
