//Take the code from the previous exercise, then store the values of type person in a map with the key of last name. Access each value in the map. Print out the values, ranging over the slice.

package main

import "fmt"

func main() {

	type person struct {
		firstName string
		lastName  string
		iceCream  string
	}

	m := map[string]person{
		"BabyBoss": person{
			firstName: "Baby",
			lastName:  "Boss",
			iceCream:  "Chocky",
		},
		"CutiePie": person{
			firstName: "Cutie",
			lastName:  "Pie",
			iceCream:  "FrenchVanilla",
		},
	}

	//fmt.Println(p1, person{})

	for k, v := range m {
		fmt.Printf("For the record %v\n", k)
		fmt.Printf("\t %v \n \t %v \n \t %v\n", v.firstName, v.lastName, v.iceCream)
	}

}
