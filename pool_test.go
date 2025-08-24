package main

import (
	"fmt"
	"sync"
	"testing"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() any {
			return "New"
		},
	}
	pool.Put("Ucup")
	pool.Put("Nur")
	pool.Put("Wahid")

	wg := sync.WaitGroup{}
	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			data := pool.Get()
			fmt.Println("data:", data)
			pool.Put(data)
		}()
	}
	wg.Wait()

	fmt.Println("Selesai")
}
