//Using the code from the previous example(code32), add a record to your map. Now print the map out using the “range” loop
//	`bond_james`, `Shaken, not stirred`, `Martinis`, `Women`
//`moneypenny_miss`, `James Bond`, `Literature`, `Computer Science`
//`no_dr`, `Being evil`, `Ice cream`, `Sunsets`

package main

import "fmt"

func main() {

	last_first := map[string][]string{
		`bond_james`:      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	fmt.Println(last_first)
	last_first["amar"] = []string{"Superpower", "Comics", "Work out"}
	fmt.Println(last_first)

	for k, v := range last_first {
		fmt.Println("This is the record for", k)

		for _, v2 := range v {
			fmt.Println("\t", v2)
		}
	}
}
