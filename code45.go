//create a type SQUARE
//create a type CIRCLE
//attach a method to each that calculates AREA and returns it
//circle area= π r 2
//square area = L * W
//create a type SHAPE that defines an interface as anything that has the AREA method
//create a func INFO which takes type shape and then prints the area
//create a value of type square
//create a value of type circle
//use func info to print the area of square
//use func info to print the area of circle
package main

import "fmt"
import "math"

type CIRCLE struct {
	radius float64
}
type SQUARE struct {
	length float64
	width  float64
}

func (c CIRCLE) AREA() float64 {
	return (c.radius * c.radius) * math.Pi
}

func (s SQUARE) AREA() float64 {
	return s.length * s.width
}

type SHAPE interface {
	AREA() float64
}

func INFO(f SHAPE) {
	fmt.Println(f.AREA())
}

func main() {

	c := CIRCLE{
		radius: 4,
	}

	s := SQUARE{
		length: 4,
		width:  10,
	}

	INFO(c)
	INFO(s)

}
