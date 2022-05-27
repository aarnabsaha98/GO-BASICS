package main

import "fmt"

func main() {
	// maps
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"California": 1,
		"Texus":      1,
		"Florida":    1,
	}
	fmt.Println(statePopulation)

	// manipulation
	fmt.Println(statePopulation["Texus"])

	statePopulation["WestBengal"] = 1
	statePopulation["Amt"] = 1
	fmt.Println(statePopulation)

	delete(statePopulation, "Amt")
	fmt.Println(statePopulation)

	// check
	pop, ok := statePopulation["Florid"]
	fmt.Println(pop, ok)

	fmt.Println(len(statePopulation))

}
