package main

import "fmt"

func main() {
	/*
		syntax of for loop with basic 3 statuements
		for statement1; statement2; stattement3 {

		}
		statement1 = initialize the loop counter
		statement2 = evalutate the loop counter
		statement3 = increase, decrease counter or control the counter
	*/

	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()
	for i := 1; i <= 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

}
