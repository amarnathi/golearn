//Create a for loop using this syntax
//for condition { }
//Have it print out the years you have been alive.

package main

import (
	"fmt"
	"time"
)

func main() {

	year := 1986

	for year <= time.Now().Year() {
		fmt.Println(year)
		year++
	}
}
