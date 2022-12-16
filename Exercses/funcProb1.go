/*
sum is a function which takes a slice of numbers
and adds them together. What would its function signature look like in Go?
*/
package main

import "fmt"

func sum(x []int) int {
	result := 0
	for _, v := range x {
		result += v
	}
	return result
}

func main() {
	slice1 := []int{2, 3, 4, 5}
	result := sum(slice1)
	fmt.Printf("The sum of the slice is %v\n", result)
}
