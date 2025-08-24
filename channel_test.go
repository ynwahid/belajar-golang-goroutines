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

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	// FIFO
	// channel <- "Ucup"
	// channel <- "Nur "
	// channel <- "Wahid"
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	go func() {
		channel <- "Ucup"
		channel <- "Nur "
		channel <- "Wahid"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println("Selesai")
	time.Sleep(2 * time.Second)
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := range 100 {
			channel <- fmt.Sprintf("Perulangan ke-%d", i)
		}
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data:", data)
	}
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	counter := 0
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer func() {
		close(channel1)
		close(channel2)
	}()

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for {
		select {
		case data1 := <-channel1:
			fmt.Println("Data dari channel 1", data1)
			counter++
		case data2 := <-channel2:
			fmt.Println("Data dari channel 2", data2)
			counter++
		}
		if counter == 2 {
			break
		}
	}
}

func TestSelectDefaultChannel(t *testing.T) {
	counter := 0
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer func() {
		close(channel1)
		close(channel2)
	}()

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	for {
		select {
		case data1 := <-channel1:
			fmt.Println("Data dari channel 1", data1)
			counter++
		case data2 := <-channel2:
			fmt.Println("Data dari channel 2", data2)
			counter++
		default:
			fmt.Println("Menunggu data...")
		}
		if counter == 2 {
			break
		}
	}
}
