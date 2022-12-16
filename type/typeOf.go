package main

import (
	"fmt"
	"reflect"
)

func main() {
	type Ishan int

	a := 1
	b := 2
	c := "Shrestha"
	d := "Tamang"
	e := []string{"Ishan", "Binod"}
	f := [...]int{1, 2, 3, 4}
	var g Ishan = 98
	fmt.Printf("Type of value a is %T\n", a)
	fmt.Println("Type of value of a using reflect package is : ", reflect.TypeOf(a))
	fmt.Printf("Type of value of b is %T\n", b)
	fmt.Printf("Type of value of c is %T\n", c)
	fmt.Printf("Type of value of d is %T\n", d)
	fmt.Printf("Type of value of e is %T\n", e)
	fmt.Printf("Type of value of f is %T\n", f)
	fmt.Printf("Type of value of g is %T\n", g)
}
