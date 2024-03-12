package main

import (
	"testing"
)

func ConcatV1() {
	var s string
	for i := 0; i < 10000; i++ {
		s += "a"
	}
}

func BenchmarkConcatV1(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ConcatV1()
	}
}
