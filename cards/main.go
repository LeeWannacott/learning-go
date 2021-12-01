package main

func main() {
	// var card string = "Ace of spades"
	// card := "Ace of Spades"
	// card = "Five of Diamonds"
	// card := newCard()
	// cards := deck{"Ace of Spades", newCard()}
	// cards = append(cards, "Six of Spades")
	// cards.print()
	// fmt.Println(cards)
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// remainingCards.print()
	// cards.print()
	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// fmt.Println(cards.toString())
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// for i, card := range cards {
	// fmt.Println(i, card)
	// }
}

// func newCard() string {
// return "Five of Diamonds"
// }
