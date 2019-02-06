//Using a COMPOSITE LITERAL:
//create a SLICE of TYPE int
//assign 10 VALUES
//Range over the slice and print the values out.
//Using format printing
//print out the TYPE of the slice

package main

import "fmt"

func main() {

	x25 := []int{
		11, 12, 21, 121, 134, 145, 156, 178, 200, 210,
	}

	for i, v := range x25 {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x25)
}
