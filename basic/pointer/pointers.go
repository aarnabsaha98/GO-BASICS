package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701
	k := &i
	fmt.Println(i, j, k)
	*k = 10

	fmt.Println(i, j, k)

	var a int = 41
	var b *int = &a

	fmt.Println(&a, b)

}
