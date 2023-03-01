package main

import (
	"fmt"
	"time"
)

func printSometing(s string) {
	fmt.Println(s)
}

func main() {
	go printSometing("THis is the first line")
	time.Sleep(1 * time.Second) // this is a very bad example
	printSometing("This is the second line")
}
