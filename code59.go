//in addition to the main goroutine, launch two additional goroutines
//each additional goroutine should print something out
//use waitgroups to make sure each goroutine finishes before your program exists

package main

import (
	"fmt"
	"sync"
)

var wg59 sync.WaitGroup

/*func foo59() {
	fmt.Println("Inside foo")
	wg59.Done()
}

func bar59() {
	fmt.Println("Inside bar")
	wg59.Done()
}*/

func main() {
	wg59.Add(2)
	//go foo59()
	//go bar59()
	go func() {
		for i59 := 0; i59 < 10; i59++ {
			fmt.Println("Inside first go func", i59)
		}
		defer wg59.Done()
	}()

	go func() {
		for j59 := 0; j59 < 10; j59++ {
			fmt.Println("Inside second go func", j59)
		}
		defer wg59.Done()
	}()

	fmt.Println("Inside Main")
	wg59.Wait()
}
