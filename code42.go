//create a func with the identifier foo that
//takes in a variadic parameter of type int
//pass in a value of type []int into your func (unfurl the []int)
//returns the sum of all values of type int passed in
//create a func with the identifier bar that
//takes in a parameter of type []int
//returns the sum of all values of type int passed in

package main

import "fmt"

func main() {
	inputParameters42 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo42(inputParameters42...))
}

func foo42(vi ...int) int {

	total := 0

	for _, v := range vi {
		total += v
	}

	return total
}
