package main

import "fmt"

func main() {
	fmt.Println("Enter the week number of choice")
	var weekNum int
	fmt.Scan(&weekNum)

	switch weekNum {
	case 1:
		fmt.Println("Sunday")
	case 2:
		fmt.Println("Monday")
	case 3:
		fmt.Println("Tuesday")
	case 4:
		fmt.Println("Wed")
	case 5:
		fmt.Println("Thu")
	case 6:
		fmt.Println("Fir")
	case 7:
		fmt.Println("Sat")
	default:
		fmt.Println("Invalid week number")
	}
}
