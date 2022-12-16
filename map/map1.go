package main

import "fmt"

func main() {
	/*
		Syntax for map
		var a = map[Key->DataType]ValueType{key1:value1, key2:value2, ...}
		a := map[Key->DataType]ValueType{key1:value1, key2:value2, ...}

	*/
	a := map[string]string{"fname": "roshan", "lname": "Ghimerer"}
	fmt.Println(a["fname"])
	/*
		Using make function to create empty map and assign to variable
	*/
	b := make(map[int]string)
	b[0] = "row"
	b[1] = "column"
	fmt.Println(b)
	fmt.Println(b[0])
	fmt.Println(b[1])
	for k, v := range a {
		fmt.Printf("%v------>%v\n", k, v)
	}
}
