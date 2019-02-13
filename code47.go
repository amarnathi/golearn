//Assign a func to a variable, then call that func

package main

import "fmt"

func main() {

	x := func() int {
		return 100
	}
	fmt.Println(x())
	foo()
}

func foo() {
	fmt.Println("Baby Boss")
}
