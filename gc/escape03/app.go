package main

func main() {
	val := 10
	inc(&val)
	println(val)
}

func inc(val *int) {
	*val++
}
