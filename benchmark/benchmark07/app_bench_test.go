package main

import (
	"math/rand"
	"testing"
	"time"
)

type MyStruct struct {
	Data [10000]int
}

func (ls MyStruct) sumValue() int {
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

func populateData(ls *MyStruct) {
	for i := range ls.Data {
		ls.Data[i] = rand.Intn(100)
	}
}

func BenchmarkSum(b *testing.B) {
	myStruct := MyStruct{}
	populateData(&myStruct)

	for i := 0; i < b.N; i++ {
		myStruct.sumValue()
	}
}
