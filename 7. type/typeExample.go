package main

import "fmt"

func main() {
	type celcius int
	type fahrenheit int

	var f fahrenheit = 32
	var c celcius = 0
	var FtoC celcius = celcius((f - 32) * 5 / 9)
	fmt.Println(f, c)
	fmt.Println(FtoC)
}
