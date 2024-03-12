package main

func main() {
	val := 10
	println(*double(val))
}

func double(val int) *int {
	res := val + val
	return &res
}
