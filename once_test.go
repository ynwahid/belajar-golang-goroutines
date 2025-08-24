package main

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnce(t *testing.T) {
	once := sync.Once{}
	wg := sync.WaitGroup{}

	for range 100 {
		go func() {
			wg.Add(1)
			once.Do(OnlyOnce)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("counter:", counter)
}
