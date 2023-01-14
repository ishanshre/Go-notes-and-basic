package main

import "fmt"

func infinite(integers ...int) {
	for _, integer := range integers {
		fmt.Println(integer)
	}
}

func main() {
	infinite(4, 2, 45, 56, 2, 45)
}
