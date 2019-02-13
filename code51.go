//Create a value and assign it to a variable.
//Print the address of that value.

package main

import (
	"fmt"
)

func main() {

	BabyBoss := 4
	fmt.Printf("%T\n", BabyBoss)
	fmt.Printf("%v\n", BabyBoss)
	fmt.Printf("%T\n", &BabyBoss)
	fmt.Printf("%v\n", &BabyBoss)

}
