package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGoMaxProcs(t *testing.T) {
	wg := sync.WaitGroup{}
	for range 100 {
		wg.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			wg.Done()
		}()
	}

	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU:", totalCPU)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread:", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total Goroutine:", totalGoroutine)

	wg.Wait()
}
