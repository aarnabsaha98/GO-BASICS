package main

import (
	"fmt"
)

// global variable

var (
	name  string  = "arnab"
	age   int     = 23
	marks float32 = 98.99
)

// FunctionName(variable  type) returnType
func main() {

	// local variable

	var i float64 = 56.09
	fmt.Printf("%v , %T", i, i)
	fmt.Println()
	age := 78.09
	fmt.Printf("%v , %T", age, age)

}
