package main

import "fmt"

type Doctor struct {
	number    int
	actorName string
	companies []string
}

func main() {
	aDoctor := Doctor{
		number:    1,
		actorName: "Arnab",
		companies: []string{
			"leo",
			"Neymar",
		},
	}

	fmt.Println(aDoctor.companies[1])
}
