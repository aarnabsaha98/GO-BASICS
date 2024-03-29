package main

import (
	"fmt"
)

func main() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(x, y float64) (float64, error) {
	if y == 0.0 {
		return 0.0, fmt.Errorf("can not divide by zero")
	}
	return x / y, nil
}
