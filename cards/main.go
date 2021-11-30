package main

import "fmt"

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()
	// fmt.Println(cards)

	cards := newDeck()
	cards.print()

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

// func newCard() string {
// return "Five of Diamonds"
// }
