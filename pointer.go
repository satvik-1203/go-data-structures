package main

func point(a int) {
	a = 20
}

func main() {
	x := 10
	// z := 11
	y := &x
	// y = &z
	point(x)
	print(x, " ", *y)
}
