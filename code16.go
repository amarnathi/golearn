//Create a for loop using this syntax
//for { }
//Have it print out the years you have been alive.

package main

import (
	"fmt"
	"time"
)

func main() {
	year16 := 1986
	for {
		if year16 > time.Now().Year() {
			break
		}
		fmt.Println("years alive are :", year16)
		year16++
	}

}
