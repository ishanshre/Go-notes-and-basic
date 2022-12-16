// panic stops the execution with displaying our message
// it is same as rasing exception using raise in python
// panic is a progammer error such as accessing index of array that does not exists
package main

import "fmt"

func main() {
	fmt.Println("First sateemt")
	panic("This is a panic and execution of the programs stop here")
	fmt.Println("after panic")
}
