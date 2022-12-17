package main

import "fmt"

func fact(x int) (y int) {
	if x > 0 {
		y = x * fact(x-1)
	} else {
		y = 1
	}
	return
}

func main() {
	z := 4
	fmt.Printf("The fact of %v is %v.", z, fact(z))
}
