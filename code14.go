//Print every rune code point of the uppercase alphabet three times. Your output should look like this:

package main

import "fmt"

func main() {

	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
