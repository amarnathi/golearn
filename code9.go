package main

import "fmt"

func main() {
	x9 := 911
	fmt.Printf("Decimal value %d\n", x9)
	fmt.Printf("Binary value %b\n", x9)
	fmt.Printf("Hexadecimal value %#X\n", x9)
	y9 := x9 << 1 //bit shift happens on binary numbers
	fmt.Printf("Decimal value of shifted bit is %d\n", y9)
	fmt.Printf("Binary value of shifted bit is %b\n", y9)
	fmt.Printf("Hexadecimal value of shifted bit is %#X\n", y9)
}
