//Building on the previous hands-on exercise, create a program that uses “else if” and “else”.
package main

import "fmt"

func main() {

	x20 := "James Bond"

	if x20 == "JAMES BOND" {
		fmt.Println("Inside If statement")
	} else if x20 == "James Bond" {
		fmt.Println("Inside else if condition")
	} else {
		fmt.Println("Inside else condition")
	}
}
