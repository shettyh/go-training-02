package main

import (
	"fmt"
	"regexp"
	"testing"
)

// Define a global compiled regular expression
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

// Function using the globally compiled regex
func isValidEmailGlobal(s string) bool {
	return emailRegex.MatchString(s)
}

// Function using regex compilation without global variable
func isValidEmailLocal(s string) bool {
	// Compile regex locally
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(s)
}

// Benchmark function for globally compiled regex
func BenchmarkIsValidEmailGlobal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidEmailGlobal("user@example.com")
	}
}

// Benchmark function for locally compiled regex
func BenchmarkIsValidEmailLocal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isValidEmailLocal("user@example.com")
	}
}

func main() {
	// Run benchmarks
	fmt.Println("Benchmarking isValidEmailGlobal:")
	globalResult := testing.Benchmark(BenchmarkIsValidEmailGlobal)
	fmt.Println(globalResult)

	fmt.Println("Benchmarking isValidEmailLocal:")
	localResult := testing.Benchmark(BenchmarkIsValidEmailLocal)
	fmt.Println(localResult)
}
