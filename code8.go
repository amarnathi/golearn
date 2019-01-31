package main

import "fmt"

//var c int = 11
//var d int

const(
	a int= 42
	b // b gets the value from a
	c =43

)


func main() {

	fmt.Printf("Typed Constant %T\t%v\n",a,a)
	fmt.Printf("Untyped Constant getting value from a %T\t%v\n",b,b)
	fmt.Printf("Untyped Constant %T\t%v\n",c,c)
	
}
