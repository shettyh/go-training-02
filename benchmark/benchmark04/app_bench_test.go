package main

import (
	"testing"
)

type MyData struct {
	Data []byte
}

func createMyData() *MyData {
	return &MyData{
		Data: make([]byte, 1024*1024), // 1MB data
	}
}

func BenchmarkCreateData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		obj := createMyData()
		_ = obj // simulate long-lived object
	}
}
