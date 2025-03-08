package main

func main() {
	x := 2
	square1(x)
	square2(x)
	square3(&x)
	square4(&x)
}

func square1(x int) int {
	return x * x
}

func square2(x int) *int {
	y := x * x
	return &y
}

func square3(x *int) {
	*x = *x * *x
}

func square4(x *int) *int {
	y := *x * *x
	return &y
}
