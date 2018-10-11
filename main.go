package main

import (
	"fmt"
)

type person struct {
	firstname string
	lastname  string
}

func main() {
	ankit := person{firstname: "Ankit", lastname: "Kumar"}
	fmt.Printf("%+v", ankit)
}
