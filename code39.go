// Add list of even integers passed

package main

import "fmt"

func main() {

	inputValues := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	count(sum, inputValues...)

}

func sum(xi ...int) int {
	total := 0

	for _, v := range xi {
		total += v
	}

	return total
}

func count(f func(sumI ...int) int, xi ...int) {

	evenNumbers := []int{}
	oddNumbers := []int{}

	for _, v := range xi {

		if v%2 == 0 {
			evenNumbers = append(evenNumbers, v)
		} else {
			oddNumbers = append(oddNumbers, v)
		}
	}
	fmt.Println("Sum of even numbers is \t", f(evenNumbers...))
	fmt.Println("Sum of odd numbers is \t", f(oddNumbers...))
}
