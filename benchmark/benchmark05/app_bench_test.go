package main

import (
	"sync"
	"testing"
)

type Data struct {
	mu    sync.Mutex
	value int
}

func readWithMutex(d *Data, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()

	d.mu.Lock()
	defer d.mu.Unlock()

	result <- d.value
}

func BenchmarkMutex(b *testing.B) {
	d := Data{value: 42}
	var wg sync.WaitGroup
	result := make(chan int, b.N)

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go readWithMutex(&d, &wg, result)
	}

	wg.Wait()
	close(result)
}
