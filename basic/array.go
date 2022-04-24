package main

import "fmt"

func main() {

	//grades := [...]int{45, 67, 78}
	var grades [3]int
	grades[0] = 1
	grades[2] = 1
	fmt.Printf("%v %T\n", grades, grades)
	fmt.Printf("%v %T\n", grades[1], grades[1])

	//length >> len() function

	// matrix
	var martixEle [3][3]int = [3][3]int{[3]int{3, 3, 3}, [3]int{4, 4, 4}, [3]int{5, 5, 5}}
	fmt.Println(martixEle)

	// copy of array
	var a [3]int
	b := a // else case we can do that address pointer
	//b := &a
	b[0] = 9
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Length : %v\n", len(b))
	fmt.Printf("Capacity : %v\n", cap(b))

	// slicing
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slicedArray := array[4:]
	fmt.Println(slicedArray)

	//make function generate an array
	arr := make([]int, 3, 5)
	fmt.Println(arr)
	fmt.Printf("Length : %v\n", len(arr))
	fmt.Printf("Capacity : %v\n", cap(arr))

	// dynamic  array
	dynamicArray := []int{}
	fmt.Println(dynamicArray)

	dynamicArray = append(dynamicArray, 1)
	dynamicArray = append(dynamicArray, 6, 7, 8, 9)
	fmt.Println(dynamicArray)
	fmt.Printf("Length : %v\n", len(dynamicArray))
	fmt.Printf("Capacity : %v\n", cap(dynamicArray))
	dynamicArray = append(dynamicArray, 2, 3, 4)
	fmt.Println(dynamicArray)
	fmt.Printf("Length : %v\n", len(dynamicArray))
	fmt.Printf("Capacity : %v\n", cap(dynamicArray))

	//2
	x := []int{1, 2, 3, 4, 5}
	//y := append(x[:2], x[3:]...) // y = [1,2,4,5] & x = [1,2,3,4,5,5]
	var y []int
	y = append(y, x[:3]...)
	y = append(y, x[4:]...) // y = [1,2,3,5] & x = [1,2,3,4,5]
	fmt.Println(y)
	fmt.Println(x)
}
