// Defer is a key which executes its statement at the end of the function call
// defer function call is pushed to the first index of the stack and is called at last

package main

import "fmt"

func first() {
	fmt.Println("First Function")
}

func second() {
	fmt.Println("second function")
}

func main() {
	defer second()
	first()
}
