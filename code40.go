//factorial using loop

package main

import "fmt"

func main() {
	fmt.Println(factorial(6))
}

func factorial(num int) int {

	n := 1

	//for i:=num; i > 0; i-- {
	//		n *= num
	//		num--
	//	}

	for ; num > 0; num-- {
		n *= num

	}

	return n
}
