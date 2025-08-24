package main

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
		close(done)
	}()

	for {
		select {
		case time := <-ticker.C:
			fmt.Println(time)
		case <-done:
			fmt.Println("Done")
			return
		}
	}
}

func TestTick(t *testing.T) {
	channel := time.Tick(1 * time.Second)
	done := make(chan struct{})

	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	for {
		select {
		case time := <-channel:
			fmt.Println(time)
		case <-done:
			fmt.Println("Done")
			return
		}
	}
}
