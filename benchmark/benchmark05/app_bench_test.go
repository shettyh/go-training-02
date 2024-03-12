package main

import (
	"fmt"
	"sync"
	"testing"
)

// Data structure to be accessed concurrently
type Data struct {
	mu    sync.RWMutex
	value int
}

// Function to read from the data structure using a regular Mutex
func readWithMutex(d *Data, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()

	d.mu.Lock()
	defer d.mu.Unlock()

	result <- d.value
}

// Function to read from the data structure using an RWMutex
func readWithRWMutex(d *Data, wg *sync.WaitGroup, result chan int) {
	defer wg.Done()

	d.mu.RLock()
	defer d.mu.RUnlock()

	result <- d.value
}

// Benchmark function to compare Mutex vs RWMutex performance
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

// Benchmark function to compare Mutex vs RWMutex performance
func BenchmarkRWMutex(b *testing.B) {
	d := Data{value: 42}
	var wg sync.WaitGroup
	result := make(chan int, b.N)

	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go readWithRWMutex(&d, &wg, result)
	}

	wg.Wait()
	close(result)
}

func main() {
	// Run benchmarks
	fmt.Println("Benchmarking Mutex:")
	mutexResult := testing.Benchmark(BenchmarkMutex)
	fmt.Println(mutexResult)

	fmt.Println("Benchmarking RWMutex:")
	rwMutexResult := testing.Benchmark(BenchmarkRWMutex)
	fmt.Println(rwMutexResult)
}
