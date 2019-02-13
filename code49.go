//A “callback” is when we pass a func into a func as an argument. For this exercise,
//pass a func into a func as an argument

package main

import "fmt"

func main() {

	bar49(foo49)
}

func foo49() int {
	return 100
}

func bar49(f func() int) {
	fmt.Println("Inside bar")
	x49 := f
	fmt.Printf("%T\n", x49)
	fmt.Println(x49())
}
