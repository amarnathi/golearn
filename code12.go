package main

import "fmt"

var a12 int

func main() {

	const (
		v12 = iota + 2019
		x12
		y12
		z12
	)

	fmt.Printf("%v\t%v\t%v\t%v\n", v12, x12, y12, z12)
	fmt.Printf("Test int type %T\n", a12)
}
