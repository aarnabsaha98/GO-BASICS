package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome , Arnab")

	var fruitList = []string{}
	fmt.Printf("type of fruitlist is %T\n ", fruitList)

	fruitList = append(fruitList, "mango", "apple", "banana", "choco", "peneat", "orange")

	fmt.Println(fruitList)

	fruitList = append(fruitList[:2], fruitList[3:]...)
	fmt.Println(fruitList)

}
