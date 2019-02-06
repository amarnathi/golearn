//create a slice using make and increase the capacity of the slice by 2 and print the capacity of the slice after assigning 2 more values after the initial set capacity
package main

import "fmt"

func main() {

	x30 := make([]string, 50, 500)
	fmt.Println(x30)
	fmt.Println(len(x30))
	fmt.Println(cap(x30))

	for i := 1; i <= 49; i++ {
		x30[i] = fmt.Sprintf("Position %d", i)
	}
	fmt.Println(x30)
	fmt.Println(len(x30))
	fmt.Println(cap(x30))

	x30 = append(x30, "Position 50", "Position 51")
	fmt.Println(x30)
	fmt.Println(len(x30))
	fmt.Println(cap(x30))

	for i := 1; i < len(x30); i++ {
		fmt.Println(i, x30[i])
	}
}
