//Write down what these print:
//fmt.Println(true && true)
//fmt.Println(true && false)
//fmt.Println(true || true)
//fmt.Println(true || false)
//fmt.Println(!true)
package main

import "fmt"

func main() {

	fmt.Println("true and true evaluates to", true && true)
	fmt.Println("true and false evaluates to", true && false)
	fmt.Println("true or true evaluates to", true || true)
	fmt.Println("true or false evaluates to", true || false)
	fmt.Println("not true evaluates to ", !true)
	fmt.Println("not false evaluates to ", !false)

}
