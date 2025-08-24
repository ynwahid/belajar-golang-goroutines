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
