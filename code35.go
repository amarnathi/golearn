//Create your own type “person” which will have an underlying type of “struct” so that it can store the following data:
//first name
//last name
//favorite ice cream flavors
//Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice which stores the favorite flavors.

package main

import "fmt"

func main() {

	type person struct {
		firstName string
		lastName  string
		iceCream  string
	}

	p1 := person{
		firstName: "Baby",
		lastName:  "Boss",
		iceCream:  "Chocky",
	}

	p2 := person{
		firstName: "Cutie",
		lastName:  "Pie",
		iceCream:  "FrenchVanilla",
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.firstName, p1.lastName, p1.iceCream)
	fmt.Println(p2.firstName, p2.lastName, p2.iceCream)
}
