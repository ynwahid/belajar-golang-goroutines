package main

import (
	"fmt"
	"sync"
	"testing"
)

func AddToMap(data *sync.Map, wg *sync.WaitGroup, value int) {
	defer wg.Done()

	data.Store(value, value)
}

func TestMap(t *testing.T) {
	data := &sync.Map{}
	wg := &sync.WaitGroup{}

	for i := range 100 {
		wg.Add(1)
		go AddToMap(data, wg, i)
	}
	wg.Wait()

	data.Range(func(key, value any) bool {
		fmt.Printf("key: %v, value: %v\n", key, value)
		return true
	})
}
