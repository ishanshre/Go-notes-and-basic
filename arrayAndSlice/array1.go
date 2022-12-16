package main

import "fmt"

func main() {
	/*
		Syntax for Array
		1. With defining length
		variableName := [length]datatype{values}
		2. Not defining length, compiler sets the length as the number of values (length is inferred)
		variableName := [...]datatype{values}

	*/
	var array = [5]bool{true, false, false, true, false}
	array1 := [4]string{"a", "b", "C", "d"}
	araay2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array1)
	fmt.Println(araay2)
	fmt.Println(array)

	/*
		access the array we use square brackets and pass index
		index starts from 0. Same as c, python or java
		array := [...]string{"hello","world"}
		fmt.Println(array[0])
		fmt.Println(array[1])
	*/
	fmt.Println(array1[0])
	fmt.Println(array1[1])
	fmt.Println(array1[2])
	fmt.Println(array1[3])

	// changing the value of array using index
	array1[0] = "xyz"
	fmt.Println(array1)

	// Initialization of array if null not initialized then it sets defaults value
	array3 := [4]int{}           // Not initialized
	array4 := [4]int{1, 2}       // partially initialized
	array5 := [4]int{1, 2, 3, 4} // fully initialized
	fmt.Println(array3)
	fmt.Println(array4)
	fmt.Println(array5)
}
