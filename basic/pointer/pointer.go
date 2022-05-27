package main

import "fmt"

func zero(x *int) {
	*x = 1
}
func main() {
	// new takes a type as an argument,
	// allocates enough memory to fit a value of that type and returns a pointer to it

	x := new(int)
	zero(x)
	fmt.Println(*x)
}
