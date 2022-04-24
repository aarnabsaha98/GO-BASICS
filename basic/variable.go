package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 45
	fmt.Printf("%v , %T", i, i)
	fmt.Println()
	var j string
	// j = string(i)  // ascii value
	// convert b/w number and string
	j = strconv.Itoa(i)
	fmt.Printf("%v , %T", j, j)
}

// cant redeclare variables but can shadow then
// lower case first letter for package scope
//upper case first letter to export
// no private scope
// no private scope
