package main

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// LargeStruct represents a large data structure
type LargeStruct struct {
	Data [10000]int
}

// Method with value receiver
func (ls LargeStruct) sumValue() int {
	sum := 0
	for _, v := range ls.Data {
		sum += v
	}
	return sum
}

// Method with pointer receiver
func (ls *LargeStruct) sumPointer() int {
	sum := 0
	for _, v := range ls.Data {
		sum += v
	}
	return sum
}

func init() {
	// Seed the random number generator for populating data
	rand.Seed(time.Now().UnixNano())
}

// Function to populate the Data field with random values
func populateData(ls *LargeStruct) {
	for i := range ls.Data {
		ls.Data[i] = rand.Intn(100)
	}
}

// Benchmark function for method with value receiver
func BenchmarkValueReceiver(b *testing.B) {
	largeStruct := LargeStruct{}
	populateData(&largeStruct)

	for i := 0; i < b.N; i++ {
		largeStruct.sumValue()
	}
}

// Benchmark function for method with pointer receiver
func BenchmarkPointerReceiver(b *testing.B) {
	largeStruct := LargeStruct{}
	populateData(&largeStruct)

	for i := 0; i < b.N; i++ {
		(&largeStruct).sumPointer()
	}
}

func main() {
	// Run benchmarks
	fmt.Println("Benchmarking ValueReceiver:")
	valueReceiverResult := testing.Benchmark(BenchmarkValueReceiver)
	fmt.Println(valueReceiverResult)

	fmt.Println("Benchmarking PointerReceiver:")
	pointerReceiverResult := testing.Benchmark(BenchmarkPointerReceiver)
	fmt.Println(pointerReceiverResult)
}
