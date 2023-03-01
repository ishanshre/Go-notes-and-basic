package main

import (
	"fmt"
	"sync"
)

func printSometing(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(s)
}

func main() {
	var wg sync.WaitGroup
	words := []string{
		"alpha",
		"beta",
		"delta",
		"gamma",
		"pi",
		"zeta",
		"eta",
		"theta",
		"epsion",
	}
	// wg.Add(9)
	wg.Add(len(words))
	for i, x := range words {
		go printSometing(fmt.Sprintf("%d: %s", i, x), &wg)
	}
	wg.Wait()
	// time.Sleep(1 * time.Second) // this is a very bad example to use for go concurency

	wg.Add(1)
	printSometing("This is the second thing to printed!", &wg)
}
