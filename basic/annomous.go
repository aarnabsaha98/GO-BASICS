package main

import "fmt"

func main() {
	//Annonomous function
	for i := 0; i < 6; i++ {
		func(i int) {
			msg := "Index"
			fmt.Println(msg, i)
		}(i) // invoking function
	}

	// we can assign a variable for a func
	f := func() {
		fmt.Println("hello")
	}
	f()

}
