package main

import "fmt"

func main() {

	Developer := struct{ name string }{name: "Arnab"}
	fmt.Println(Developer.name)

	// use pointer to change on the source
	// as soon as the copy change
	QualityAssurance := Developer

	QualityAssurance.name = "Saha"
	fmt.Println(QualityAssurance.name)

}
