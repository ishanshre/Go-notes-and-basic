/*
Write a function which takes an integer and
halves it and returns true if it was even or false
if it was odd. For example half(1) should return
(0, false) and half(2) should return (1,
true).
*/

package main

import "fmt"

func half(x int) (val int, status bool) {
	halfed := x / 2
	if x%2 == 0 {
		status := true
		return halfed, status
	} else {
		status := false
		return halfed, status
	}

}

func main() {
	var num int
	fmt.Println("Enter a number: ")
	fmt.Scan(&num)
	halfed, status := half(num)
	fmt.Printf("Halfed Vlaue is %v and status is %t", halfed, status)

}
