package main

import "fmt"

func main() {
	// slice does not have the length limit like the array.
	// len() return the size of the array
	// cap() returns the capacity of array upto which it can grow and shrink to
	slice1 := []string{"a", "b"}
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
	slice1 = append(slice1, "c", "d")
	fmt.Println(slice1)
	slice2 := []string{"Hello", "Name"}
	slice3 := append(slice1, slice2...)
	fmt.Println(slice3)
	slice4 := []int{2, 3, 4, 55, 6}
	fmt.Println("The slice 4 is: ", slice4)
	fmt.Println("Now appending int to the slice 4")
	slice4 = append(slice4, 1, 2, 3, 4, 4)
	fmt.Println("The appended slice4 is ", slice4)
}
