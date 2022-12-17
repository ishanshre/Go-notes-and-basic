package main

import "fmt"

type detail struct {
	firstName string
	lastName  string
	age       int
}

func (d detail) fullName() {
	fmt.Println(d.firstName + " " + d.lastName)
}

func main() {
	binod := detail{
		firstName: "Binod",
		lastName:  "Waiba Tamang",
		age:       21,
	}
	binod.fullName()
}
