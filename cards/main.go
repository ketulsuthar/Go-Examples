package main

import "log"

func main() {
	// Get new deck
	cards := newDeck()

	cards.print()

	// Get deal
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()

	// Save decks to file
	err := cards.saveToFile("my_cards")
	if err != nil {
		log.Fatal(err)
	}

	// Load decks from saved file
	cards = newDeckFromFile("my")
	// print decks
	cards.print()

	// shuffle decks
	cards.shuffle()
	cards.print()
}
