package main

import (
	"fmt"
	"strings"
)

func main() {
	// for upper case strings.ToUpper(a)  ---- for lower strings.ToLower(varName)
	a := "UndersStasdnbadaAsdf"
	fmt.Println(strings.ToUpper(a))
	fmt.Println(strings.ToLower(a))
	// strings.HasPrefix ---> searches the string from the beginning
	// strings.HasSuffix ----> searches the string from the end
	// strings.Contains ----> searches any where in the string
	// strings.Count ----> Count how many times the string appears in the string
	b := "Hello world"
	fmt.Println(strings.HasSuffix(b, "world"))
	fmt.Println(strings.HasPrefix(b, "Hello"))
	fmt.Println(strings.Contains(b, "w"))
	fmt.Println(strings.Contains(b, "w "))
	fmt.Println(strings.Count(b, "l"))
	c := []string{"sharks", "crabs", "whales"}
	fmt.Println(strings.Join(c, "_"))
	fmt.Println(strings.Join(c, ", "))
}
