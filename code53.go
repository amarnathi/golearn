//Hands-on exercise #1
//Starting with this code, marshal the []user to JSON. There is a little bit of a curve ball in this hands-on exercise - remember to ask yourself what you need to do to EXPORT a value from a package.

package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		"Baby",
		4,
	}

	u2 := user{
		"Cutie",
		2,
	}

	users := []user{u1, u2}

	fmt.Println(users)

	b, _ := json.Marshal(u1)
	fmt.Println(string(b))

}
