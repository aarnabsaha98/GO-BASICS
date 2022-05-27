package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("The test")
	}
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"California": 1,
		"Texus":      1,
		"Florida":    1,
	}

	if res, ok := statePopulation["Florid"]; ok {
		fmt.Println(res)
	} else {
		fmt.Println("No Values Found")
	}

	// if else structure

	number := 9
	guess := 3
	if guess < 1 {

	} else if guess > 10 {

	} else {
		if number == guess {

		} else if number > guess {

		} else {

		}
	}
	myNum := 0.123
	if myNum == math.Pow(math.Sqrt(myNum), 2) {
		fmt.Println("Same")
	} else {
		fmt.Println("Different")
	}

	// switch case type 1
	switch 2 {
	// can pass multiple case in one
	case 1, 3, 5:
		fmt.Println("ODD")
	case 2, 4:
		fmt.Println("Even")
	default:
		fmt.Println("none is correct")
	}
	// switch case type 2
	i := 10

	switch {
	case i <= 10:
		fmt.Println()
	case i > 10:
		fmt.Println()
	default:
		fmt.Println()
	}

}
