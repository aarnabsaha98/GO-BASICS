package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	Name   string `required max :"100"`
	Origin string
}

type Bird struct {
	// Inheritance
	Animal
	speed  float32
	canFly bool
}

func main() {
	b := Bird{}
	b.Name = "CHi"
	b.Origin = "India"
	b.speed = 23
	b.canFly = true

	fmt.Println(b.Name)

	//
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag)
	fmt.Println(field)

}
