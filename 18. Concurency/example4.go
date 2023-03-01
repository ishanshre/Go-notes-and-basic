package main

import (
	"fmt"
	"sync"
)

var msg string
var wg sync.WaitGroup

func updateMessage(s string) {
	defer wg.Done()
	msg = s
}

func printMessage() {
	fmt.Println(msg)
}

func main() {
	msgs := []string{"Hello, universe", "Hello, cosmic", "Hello, world"}
	for _, v := range msgs {
		wg.Add(1)
		updateMessage(v)
		wg.Wait()
		printMessage()
	}
}
