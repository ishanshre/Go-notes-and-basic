package main

import "fmt"

func main() {
	x := 1
	y := x
	var ptr *int = &x
	fmt.Println(&x)
	fmt.Println(&y)
	fmt.Println(ptr)
}
