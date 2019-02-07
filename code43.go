//Use the “defer” keyword to show that a deferred func runs after the func containing it exits.
package main

import "fmt"

func main() {
	defer defered()
	undefered()
}

func defered() {
	fmt.Println("defered function expected to print at last , though it is called first")
}

func undefered() {
	fmt.Println("undefered function expected to print at first , though it is called last")
}
