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
	lastname  string
	contact   contactInfo
}

func main() {
	ankit := person{
		firstname: "Ankit",
		lastname:  "Kumar",
		contact: contactInfo{
			email:   "ankit@gmail.com",
			zipCode: 1232,
		},
	}

	pointerToAnkit := &ankit
	pointerToAnkit.updateName("rahul")
	fmt.Printf("%+v", ankit)
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstname = newFirstName
}
