//Create a new type: vehicle.
//The underlying type is a struct.
//The fields:
//doors
//color
//Create two new types: truck & sedan.
//The underlying type of each of these new types is a struct.
//Embed the “vehicle” type in both truck & sedan.
//Give truck the field “fourWheel” which will be set to bool.
//Give sedan the field “luxury” which will be set to bool. solution
//Using the vehicle, truck, and sedan structs:
//using a composite literal, create a value of type truck and assign values to the fields;
//using a composite literal, create a value of type sedan and assign values to the fields.
//Print out each of these values.
//Print out a single field from each of these values.

package main

import "fmt"

func main() {

	type vehicle struct {
		doors int
		color string
	}

	type truck struct {
		name string
		vehicle
		fourWheel bool
	}

	type sedan struct {
		name string
		vehicle
		luxury bool
	}

	dodgeRam := truck{
		name: "dodgeRam",
		vehicle: vehicle{
			doors: 8,
			color: "blue",
		},
		fourWheel: false,
	}

	accord := sedan{
		name: "accord",
		vehicle: vehicle{
			doors: 4,
			color: "red",
		},
		luxury: true,
	}

	fmt.Println("For the type vehicle", dodgeRam.name)
	fmt.Println("\t", dodgeRam.vehicle.doors, dodgeRam.vehicle.color, dodgeRam.fourWheel)

	fmt.Println("For the type vehicle", accord.name)
	fmt.Println("\t", accord.vehicle.doors, accord.vehicle.color, accord.luxury)

}
