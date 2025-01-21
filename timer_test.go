package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Timer started", time.Now())

	times := <-timer.C
	fmt.Println("Timer expired at", times)
}

func TestTimeAfter(t *testing.T) {
	channel := time.After(5 * time.Second)
	fmt.Println("Timer started", time.Now())

	times := <-channel
	fmt.Println("Timer expired at", times)
}

func TestTimeAfterFunc(t *testing.T) {
	wg := &sync.WaitGroup{}
	wg.Add(1)

	// use after func to delay job, make sure to use wait group to wait for afterFunc to finish
	time.AfterFunc(1*time.Second, func() {
		fmt.Println("after func", time.Now())
		wg.Done()
	})
	fmt.Println(time.Now())

	wg.Wait()
}
