package main

import "fmt"

func main() {

	for i := 33; i < 98; i++ {
		fmt.Printf("Ascii value of %v is %q\n", i, i)
	}
	fmt.Println("")
}
