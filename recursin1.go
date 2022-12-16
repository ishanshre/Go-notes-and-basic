package main

import "fmt"

func main() {
	testCount(1)
}

func testCount(x int) int {
	if x == 11 {
		return 0
	}
	fmt.Println(x)
	return testCount(x + 1)
}
