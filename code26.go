//Using the code from the previous example, use SLICING to create the following new slices which are then printed:
//[42 43 44 45 46]
//[47 48 49 50 51]
//[44 45 46 47 48]
//[43 44 45 46 47]

//slices start from 1
//if you have a slice with 4:9, meaning it starts after 4 and ends with 9th character
package main

import "fmt"

func main() {

	a26 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(a26)
	fmt.Println(a26[:5])
	fmt.Println(a26[5:])
	fmt.Println(a26[2:7])
	fmt.Println(a26[1:6]) // it starts after 1 and ends with 6 ,including 6
}
