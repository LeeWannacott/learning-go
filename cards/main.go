package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	card := newCard()
	fmt.Println(card)
}

func newCard() string {
	return "Five of Diamonds"
}
