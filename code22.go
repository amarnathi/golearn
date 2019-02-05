//Create a program that uses a switch statement with the switch expression specified as a variable of TYPE string with the IDENTIFIER “favSport”.

package main

import "fmt"

func main() {
	favSport := "cricket"
	switch favSport {
	case "football", "basketball":
		fmt.Println("football nor basketball")
	case "tennis", "shuttle":
		fmt.Println("tennis or shuttle")
	case "cricket":
		fmt.Println("cricket")
	default:
		fmt.Println("default")
	}
}
