//Starting with this code, unmarshal the JSON into a Go data structure. Hint: https://mholt.github.io/json-to-go/

package main

import (
	"encoding/json"
	"fmt"
)

type person55 struct {
	First   string   `json:"First"`
	Last    string   `json:"Last"`
	Age     int      `json:"Age"`
	Sayings []string `json:"Sayings"`
}

var persons55 []person55

func main() {
	s55 := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	fmt.Println(persons55)
	fmt.Printf("%T\t%T\n", &persons55, persons55)
	json.Unmarshal([]byte(s55), &persons55)
	fmt.Println(persons55)

}
