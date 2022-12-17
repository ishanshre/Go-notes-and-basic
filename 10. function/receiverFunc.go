package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Printf("Name --> %s\nAge --> %d", p.name, p.age)
}

func (d person) display() {
	fmt.Printf(d.name)
}

func main() {
	binod := person{
		name: "Binod Tamang",
		age:  21,
	}
	binod.display()
}
