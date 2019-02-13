// Sort the users by Last name
// Sort them by Age
// Sort them by Sayings

//Also another of doing it

// https://godoc.org/sort#example-package

package main

import (
	"fmt"
	"sort"
)

type user58 struct {
	First58   string
	Last58    string
	Age58     int
	Sayings58 []string
}

func main() {
	u158 := user58{
		First58: "James",
		Last58:  "Bond",
		Age58:   32,
		Sayings58: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u258 := user58{
		First58: "Miss",
		Last58:  "Moneypenny",
		Age58:   27,
		Sayings58: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u358 := user58{
		First58: "M",
		Last58:  "Ammmm",
		Age58:   54,
		Sayings58: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users58 := []user58{u158, u258, u358}

	//fmt.Println(users58)
	sort.SliceStable(users58, func(i, j int) bool { return users58[i].Last58 < users58[j].Last58 })
	fmt.Println("\nSorted by Last name :", users58, "\n\n")

	sort.SliceStable(users58, func(i, j int) bool { return users58[i].Age58 < users58[j].Age58 })

	fmt.Println("Sort by Age:", users58, "\n\n")

	for _, v58 := range users58 {

		sort.SliceStable(v58.Sayings58, func(i, j int) bool {
			return v58.Sayings58[i] < v58.Sayings58[j]
		})

	}
	fmt.Println("Sort by Sayings with in the users:", users58, "\n")
}
