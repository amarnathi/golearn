package main

import (
	"fmt"
	"runtime"
)

type tofu5 int
var x5 tofu5
var y5 int

func main() {
	fmt.Printf("Default x5 value is %v\n", x5)
	fmt.Printf("x5 datatype is %T\n", x5)
	x5 = 42
	fmt.Printf("x5 assignment is %v\n",x5)
	y5 = int(x5)
	fmt.Printf("y5 datatype is %T\n",y5)
	fmt.Printf("y5 value is %v\n",y5)

	abc:="2+2"
	fmt.Println(abc)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

}
