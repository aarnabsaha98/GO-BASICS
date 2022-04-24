package main

import (
	"fmt"
)

func main() {
	number := 70
	guess := -70

	if guess < 1 {
		fmt.Println("Out of Box for 1")
	} else if guess > 100 {
		fmt.Println("Out of Box for 100")
	} else {
		if number == guess {
			fmt.Println("got IT")
		}
		if guess < number {
			fmt.Println("Low")
		}
		if guess > number {
			fmt.Println("High")
		}
	}

}

func myfunc() bool {
	fmt.Println("returned True")
	return true
}
