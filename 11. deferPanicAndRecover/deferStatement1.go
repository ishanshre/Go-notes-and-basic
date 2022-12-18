package main

import "fmt"

/*
	defer is a statment that is executed until surrounding function returns
	defer functin calls are pushed into stack in last in first out principle
*/

func main() {
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done counting")
}
