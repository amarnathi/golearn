//Starting with this code, encode to JSON the []user sending the results to Stdout. Hint: you will need to use json.NewEncoder(os.Stdout).encode(v interface{})

package main

import (
	"encoding/json"
	"os"
)

type user56 struct {
	First56   string
	Last56    string
	Age56     int
	Sayings56 []string
}

func main() {
	u156 := user56{
		First56: "James",
		Last56:  "Bond",
		Age56:   32,
		Sayings56: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u256 := user56{
		First56: "Miss",
		Last56:  "Moneypenny",
		Age56:   27,
		Sayings56: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u356 := user56{
		First56: "M",
		Last56:  "Hmmmm",
		Age56:   54,
		Sayings56: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users56 := []user56{u156, u256, u356}

	//fmt.Println(users56)

	// your code goes here

	json.NewEncoder(os.Stdout).Encode(users56)
}
