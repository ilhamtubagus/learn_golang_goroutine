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
