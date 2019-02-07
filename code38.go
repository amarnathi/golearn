//Create and use an anonymous struct.

package main

import "fmt"

func main() {

	civic := struct {
		doors  int
		luxury bool
	}{
		doors:  4,
		luxury: false,
	}

	fmt.Println(civic.luxury, civic.doors)

}
