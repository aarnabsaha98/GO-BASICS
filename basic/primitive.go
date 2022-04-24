package main

import (
	"fmt"
)

func main() {
	N := 1
	fmt.Printf("%v , %T\n", N, N)

	// declare bool
	var isBool bool // Zero value False
	fmt.Printf("%v , %T\n", isBool, isBool)

	M := 1 == 3
	fmt.Printf("%v , %T\n", M, M)

	// un assigned integer
	var unInt uint16 = 42
	fmt.Printf("%v , %T\n", unInt, unInt)

	var a int16 = 9
	var b int = 5
	fmt.Println(int(a) + b)

	// bin ops
	x := 10             // 	1010
	y := 3              // 	0011
	z := 8              // 2^3
	fmt.Println(x & y)  // 0010
	fmt.Println(x | y)  // 1011
	fmt.Println(x ^ y)  // 1001
	fmt.Println(x &^ y) // 0100
	fmt.Println(z << 3) // 2^3 * 2^3 = 64
	fmt.Println(z >> 3) // 2^3 / 2^3 = 1

	// COMPLEX NUMBERS
	var comp complex128 = 3 + 2i
	var comp1 complex128 = 5 + 2i
	fmt.Printf("%v %T\n", comp, comp)
	fmt.Println(comp + comp1)
	var mul complex128 = comp + comp1
	fmt.Printf("%v %T\n", (mul), (mul))
	fmt.Printf("%v %T\n", real(mul), real(mul))
	fmt.Printf("%v %T\n", imag(mul), imag(mul))

	// CONVERT  NUMBERS INTO COMPLEX NUMBERS
	var revertComplex complex128 = complex(4, 10)
	fmt.Printf("%v %T", revertComplex, revertComplex)

}
