//Closure is when we have “enclosed” the scope of a variable in some code block. For this hands-on exercise, create a func which “encloses” the scope of a variable:

package main

import "fmt"

func main() {
	a := foo50()
	b := foo50()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())

}

func foo50() func() int {
	x := 10
	return func() int {
		x++
		return x
	}
}
