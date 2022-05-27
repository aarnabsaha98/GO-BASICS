package main

import (
	"fmt"
)

func main() {
	s := "This is a GoLang , type::"
	fmt.Printf("%v %T\n", s, s)
	// String is ImMutable
	// we need to convert it into string after accessing any
	// chars
	fmt.Printf("%v %T\n", string(s[2]), string(s[2]))
	// string holds utf8 char values
	converUTF := []byte(s)
	fmt.Printf("%v %T", converUTF, converUTF)
}
