package main

import "fmt"

func main() {
	/*
		Syntax using range keyword
		for index, value = range array|slice|map
	*/
	array1 := [4]int{1, 2, 3, 4}
	for index, value := range array1 {
		fmt.Printf("%v--->%v", index, value)
		fmt.Println()
	}
	slice1 := []string{"Hello", "world"}
	for index, value := range slice1 {
		fmt.Printf("%v---->%v", index, value)
		fmt.Println()
	}
	// if we only want index or value we use underscore to replace varialbe that we do not want
	for index, _ := range array1 {
		fmt.Println(index)
	}
	for _, value := range array1 {
		fmt.Println(value)
	}

}
