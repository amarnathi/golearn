//Using the code from the previous example, delete a record from your map. Now print the map out using the “range” loop
package main

import "fmt"

func main() {

	last_first := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(last_first)

	if _, ok := last_first["bond_james"]; ok {
		delete(last_first, "bond_james")
		fmt.Println("Deleted bond_james key")
	}
	if _, ok := last_first["amar"]; ok {
		delete(last_first, "amar")
		fmt.Println("Deleted")
	}

	fmt.Println(last_first)

}
