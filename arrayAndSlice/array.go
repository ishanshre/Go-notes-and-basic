package main

import "fmt"

var a string = "Hello world"

const (
	b int    = 10
	c string = "HEllo"
	d bool   = true
)

var (
	e = 20
)

func main() {
	cards := []string{"Five of diamond", newCard()} // creating a slice
	var card1 string = "30"
	fmt.Println(cards)
	fmt.Println(card1)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(e)
}

func newCard() string {
	return "Hellow World"
}
