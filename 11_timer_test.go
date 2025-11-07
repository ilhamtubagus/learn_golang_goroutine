package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// Create a timer that will send the current time to its channel
// time.NewTimer will return Timer instance (from struct consisting of C (channel) and initTimer)

func TestTimer(t *testing.T) {
	timer := time.NewTimer(5 * time.Second)
	fmt.Println("Timer started", time.Now())

	// Will print current time until the timer finished
	for {
		select {
		case <-timer.C:
			fmt.Println("Timer expired", time.Now())
			return
		default:
			fmt.Println("Timer running", time.Now())
		}
	}
}

// Similar to time.NewTimer but time.After directly returns Channel
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
	fmt.Println("printed first " + time.Now().String())

	wg.Wait()
}
