package main

import "fmt"

type tofu int

var x4 tofu

func main() {

	fmt.Printf("Value of x4 is %v\n", x4)
	fmt.Printf("Type of x4 value is %T\n", x4)
	x4 = 42
	fmt.Printf("Assinged value of x4 is %v\n", x4)
}
