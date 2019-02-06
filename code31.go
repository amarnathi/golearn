//Create a slice of a slice of string ([][]string). Store the following data in the multi-dimensional slice:
//
//"James", "Bond", "Shaken, not stirred"
//"Miss", "Moneypenny", "Helloooooo, James."
//
//Range over the records, then range over the data in each record.

package main

import "fmt"

func main() {

	slice1 := []string{"James", "Bond", "Shaken, not stirred"}
	slice2 := []string{"Miss", "MoneyPenny", "Hellooooo, James."}
	fmt.Println(slice1)
	fmt.Println(slice2)
	multislice := [][]string{slice1, slice2}
	fmt.Println(multislice)
	for i := 0; i < len(multislice); i++ {
		fmt.Println(i, multislice[i])
		for j := 0; j < len(multislice[i]); j++ {
			fmt.Println(j, multislice[i][j])
		}
	}
}
