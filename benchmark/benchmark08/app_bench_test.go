package main

import "testing"

var data = map[string]string{
	"hello": "world",
	"world": "hello",
	"a":     "b",
}

func CheckAndReturn(val string) string {
	if _, ok := data[val]; ok {
		return data[val]
	}
	return ""
}

func BenchmarkCheckAndReturn(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		CheckAndReturn("hello")
	}
}
