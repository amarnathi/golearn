package main

import "fmt"

var x3 int = 42
var y3 string = "James Bond"
var z3 bool = true

func main() {

	s := fmt.Sprintf("%v%v%v", x3, y3, z3)
	fmt.Println(s)

}
