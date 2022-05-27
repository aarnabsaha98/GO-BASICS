package main

import "fmt"

func main() {
	/*
			// 1 for
			for j := 0; j < 10; j++ {

			}

			// 2  D0 - While
			i := 0
			for ; i < 5; i++ {
				fmt.Println(i)
				if i%2 == 0 {
					fmt.Println("even :::", i)
				} else {
					fmt.Println("odd ::: ", i)
				}
			}

			// 3 While
			k := 1
			for k < 2 {
				fmt.Print(i)
				k++
			}

		// infinite loop

		x := 0

		for {
			fmt.Println(x)
			x++
			if x == 4 {
				break
			}
		}
	*/

	// nested loop
	/*
		Loop:
			for i := 1; i < 5; i++ {

				for j := 1; j < 3; j++ {
					fmt.Println("i , j  ::", i, j)
					fmt.Println(i * j)
					if i*j == 3 {
						fmt.Println(" hey  :: ", i*j)

						break Loop

					}
				}
			}
	*/

	// from slice/ array
	s := []int{1, 2, 3}
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}

	fmt.Println("print length s ::", len(s))

	// range ( to loop over and get index and value)

	for k, v := range s {
		fmt.Println(k, v)
	}

}
