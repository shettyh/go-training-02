package main

func main() {
	val := 10

	var iface Math
	iface = Summer{}
	println(iface.Double(val))

	//var concrete Summer
	//res := concrete.Double(val)
	//println(res)
}

type Math interface {
	Double(a int) int
}

type Summer struct {
}

func (s Summer) Double(a int) int {
	return a + a
}
