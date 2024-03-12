package main

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

const (
	start = 10  // actual = start * goprocs
	end   = 100 // actual = end    * goprocs
	step  = 20
)

var goprocs = runtime.GOMAXPROCS(0) // 8

func BenchmarkChanWrite(b *testing.B) {
	var v int64
	ch := make(chan int, 1)
	ch <- 1
	for i := start; i < end; i += step {
		b.Run(fmt.Sprintf("goroutines-%d", i*goprocs), func(b *testing.B) {
			b.SetParallelism(i)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					<-ch
					v += 1
					ch <- 1
				}
			})
		})
	}
}

func BenchmarkMutexWrite(b *testing.B) {
	var v int64
	mu := sync.Mutex{}
	for i := start; i < end; i += step {
		b.Run(fmt.Sprintf("goroutines-%d", i*goprocs), func(b *testing.B) {
			b.SetParallelism(i)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					mu.Lock()
					v += 1
					mu.Unlock()
				}
			})
		})
	}
}
