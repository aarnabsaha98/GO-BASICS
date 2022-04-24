package main

import (
	"fmt"
)

func main() {
	greetings("arnab")

	for i := 0; i < 5; i++ {
		tellYourName("TOM", i+1)
	}
	s := sum(1, 2, 3, 4, 5)

	fmt.Println("sum", s)
}

func greetings(name string) {
	fmt.Println("hello ", name, "You have done it")
}

func tellYourName(name string, idx int) {
	fmt.Println()
}

func sum(values ...int) int {
	fmt.Println(values)
	res := 0
	for _, v := range values {
		res += v

	}
	// fmt.Println("the sum is ", res)
	return res
}
