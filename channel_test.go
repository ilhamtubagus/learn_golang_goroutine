package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(1 * time.Second)
		channel <- "Hello, World!"
		fmt.Println("Data sent to channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func GiveMeResponse(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello, World!"
}

func TestChannelAsParameter(t *testing.T) {
	responseChannel := make(chan string)
	defer close(responseChannel)

	go GiveMeResponse(responseChannel)

	data := <-responseChannel
	fmt.Println(data)

	time.Sleep(2 * time.Second)
}

func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Hello, World!"
	fmt.Println("Data sent to channel")
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannelDirection(t *testing.T) {
	responseChannel := make(chan string)
	defer close(responseChannel)

	go OnlyIn(responseChannel)
	OnlyOut(responseChannel)

	time.Sleep(1 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 1)
	defer close(channel)

	go func() {
		channel <- "Hello, World!"
		channel <- "Another message"

	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 5; i++ {
			channel <- fmt.Sprintf("Message %d", i)
		}
		defer close(channel)
	}()

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("Done")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Received data from channel1:", data)
			counter++
		case data := <-channel2:
			fmt.Println("Received data from channel2:", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}

}
