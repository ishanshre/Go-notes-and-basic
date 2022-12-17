package main

import "fmt"

func main() {
	num := 9
	if num < 9 {
		fmt.Println("less than 9")
	} else {
		fmt.Println("greater or equal to 9")
	}

	if a := 10; a < 20 { // declaring and initializing in same if else statement
		fmt.Println("a is less than 20")
	} else {
		fmt.Println("a is greateer than 20")
	}
}
