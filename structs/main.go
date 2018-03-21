package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstname string
	name      string
	contact   contactInfo
}

func main() {
	//alex := person{firstname: "Alex", name: "Anderson"}

	// var alex person
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	pepito := person{
		firstname: "Pepito",
		name:      "localsonne",
		contact: contactInfo{
			email:   "flfl@ldl.Fr",
			zipCode: 92400,
		},
	}

	fmt.Printf("%+v", pepito)
}
