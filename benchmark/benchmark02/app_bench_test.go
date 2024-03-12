package main

import (
	"testing"
)

func AppendV1() {
	var slice []int
	for i := 0; i < 10000; i++ {
		slice = append(slice, i)
	}
}

func BenchmarkAppendV1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		AppendV1()
	}
}
