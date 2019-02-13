//Build and use an anonymous func

package main

import "fmt"

func main() {

	x := func() int {
		return 100
	}()

	fmt.Println(x)

}
