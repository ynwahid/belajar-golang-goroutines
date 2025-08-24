package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int64
	wg := sync.WaitGroup{}
	for i := 1; i <= 100; i++ {
		wg.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&x, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("counter:", x)
}
