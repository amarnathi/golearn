//create a person struct
//create a func called “changeMe” with a *person as a parameter
//change the value stored at the *person address
//important
//to dereference a struct, use (*value).field
//p1.first
//(*p1).first
//“As an exception, if the type of x is a named pointer type and (*x).f is a valid selector expression denoting a field (but not a method), x.f is shorthand for (*x).f.”
//https://golang.org/ref/spec#Selectors
//create a value of type person
//print out the value
//in func main
//call “changeMe”
//in func main
//print out the value

package main

import (
	"fmt"
)

type person52 struct {
	first string
	last  string
	age   int
}

func changeMe(p *person52) {

	fmt.Printf("%T\n", p.first)
	fmt.Printf("%T\n", &p.first)

	fmt.Printf("%v\n", p.first)
	fmt.Printf("%v\n", &p.first)

	p1 := &p.first
	*p1 = "Cutie"

}

func main() {

	p52 := person52{
		first: "Baby",
		last:  "Boss",
		age:   4,
	}
	fmt.Println(p52)
	changeMe(&p52)
	fmt.Println(p52)

}
