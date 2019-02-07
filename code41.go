//create a func with the identifier foo that returns an int
//create a func with the identifier bar that returns an int and a string
//call both funcs
//print out their results

package main

import "fmt"

func main() {
	fmt.Println("Output of foo() function is \t", foo41())
	x, y := bar41()
	fmt.Println(x, y)
}

func foo41() int {
	return 111
}

func bar41() (int, string) {
	return 42, "baby"
}
