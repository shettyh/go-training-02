package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	go func() {
		http.ListenAndServe(":6060", nil)
	}()

	// Simulate some work for profiling purposes
	for i := 0; i < 100000; i++ {
		doSomeWork()
	}

	// Sleep for a while to allow time for profiling
	time.Sleep(10 * time.Second)
}

func doSomeWork() {
	// Simulate CPU-intensive work
	for i := 0; i < 10000; i++ {
		_ = i * i
	}

	// Simulate I/O work
	time.Sleep(1 * time.Millisecond)
}
