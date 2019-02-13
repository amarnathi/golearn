//Create a func which returns a func
//assign the returned func to a variable
//call the returned func

package main

import "fmt"

func main() {

	x48 := foo48()

	fmt.Printf("%T\n", x48)
	fmt.Println(x48())

}

func foo48() func() int {
	return func() int {
		return 100
	}
}
