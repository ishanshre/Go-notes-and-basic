package main

import "fmt"

func main() {
	var card string = "Ace of Spades"
	card1 := 12
	card3 := "king" // := tells the compiler to figure out the type of variable it is
	card4 := newCard()
	card5 := newCard1()
	card6 := newCard2()
	fmt.Println(card)
	fmt.Println(card1)
	fmt.Println(card3)
	fmt.Println(card4)
	fmt.Println(card5)
	fmt.Println(card6)
}

func newCard() int {
	return 5
}

func newCard1() string {
	return "Return string"
}

func newCard2() bool {
	return false
}
