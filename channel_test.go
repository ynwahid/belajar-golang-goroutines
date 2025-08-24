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
		time.Sleep(2 * time.Second)
		// channel should be filled
		channel <- "Ucup"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	// and should be consumed
	data := <-channel
	fmt.Println(data)
	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Ucup"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)
	fmt.Println(<-channel)
	time.Sleep(5 * time.Second)
}

// send
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	fmt.Println("Sending Data")
	channel <- "Ucup"
}

// receive
func OnlyOut(channel <-chan string) {
	fmt.Println("Receiving Data")
	data := <-channel
	fmt.Println(data)
}

func TestChannelInChannelOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
