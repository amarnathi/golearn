//Create a program that uses a switch statement with no switch expression specified.

package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Inside false case")
	case true:
		fmt.Println("Inside true case")
	default:
		fmt.Println("Inside default case")
	}
}
