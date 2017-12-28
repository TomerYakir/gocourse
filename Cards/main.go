package main

import "fmt"

var deckSize int

func main() {
	// var card string = "Ace of Spades" // always a string
	// var card = "Ace of Spades"
	// card := "Ace of Spades"
	// card = newCard()

	// cards := []string{newCard(), newCard(), "Ace of Diamonds"} //slice of strings
	// cards := deck{newCard(), "Ace of Diamonds"}
	// cards = append(cards, "Six of Spades") // append returns a new slice
	// cards := newDeck()
	// cards.print()
	// hand, cards := cards.deal(3)
	// hand.print()
	// cards.saveToFile("all_cards")
	cards := readDeckFromFile("all_cards")
	cards.print()
	fmt.Println("shuffling...")
	cards = cards.shuffle()
	cards.print()

	/*
		fmt.Println(card)
		fmt.Println(cards)
	*/

}

func newCard() string {
	return "Five of Diamonds"
}
