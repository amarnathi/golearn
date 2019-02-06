//Using a COMPOSITE LITERAL:
//create an ARRAY which holds 5 VALUES of TYPE int
//assign VALUES to each index position.
//Range over the array and print the values out.
//Using format printing
//print out the TYPE of the array

package main

import "fmt"

func main() {

	x24 := [5]int{}
	x24[0] = 00
	x24[1] = 01
	x24[2] = 02
	x24[3] = 03
	x24[4] = 04

	for i, v := range x24 {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x24)

	//for i := 0; i < len(x24); i++ {
	//	fmt.Printf("%T\t%v\n", i, i)
	//}
}
