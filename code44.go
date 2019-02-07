//Create a user defined struct with
//the identifier “person”
//the fields:
//first
//last
//age
//attach a method to type person with
//the identifier “speak”
//the method should have the person say their name and age
//create a value of type person
//call the method from the value of type person

package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "Baby",
		last:  "Boss",
		age:   4,
	}

	p2 := person{
		first: "Cutie",
		last:  "Pie",
		age:   2,
	}
	p1.speak()
	p2.speak()
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}
