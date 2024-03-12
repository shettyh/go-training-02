package main

import (
	"fmt"
	"sync"
)

// SharedData represents a structure with a shared counter
type SharedData struct {
	counter int
}

// IncrementCounter increments the counter in SharedData without proper synchronization
func (sd *SharedData) IncrementCounter() {
	// Simulate some non-atomic operation
	current := sd.counter
	sd.counter = current + 1
}

// DisplayCounter displays the current counter value in SharedData
func (sd *SharedData) DisplayCounter() {
	fmt.Printf("Counter: %d\n", sd.counter)
}

func main() {
	// Create a shared data structure
	sharedData := SharedData{}

	// Number of goroutines to simulate concurrent access
	numGoroutines := 10

	// Wait group to wait for all goroutines to finish
	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Simulate concurrent access to the shared data without proper synchronization
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			sharedData.IncrementCounter()
			sharedData.DisplayCounter()
		}()
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
