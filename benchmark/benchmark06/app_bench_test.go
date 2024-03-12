package main

import (
	"regexp"
	"testing"
)

func isValidEmailLocal(s string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(s)
}

func BenchmarkIsValidEmailLocal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidEmailLocal("user@example.com")
	}
}
