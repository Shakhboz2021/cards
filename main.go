package main

func main() {
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// fmt.Println("------------")
	// remainingCards.print()

	// cards := newDeck()
	// cards.saveToFiles("my_cards")
	// cardsFromFile := newDeckFromFile("my_cards")
	// cardsFromFile.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()

}

func newCard() string {
	return "Five of Spades"
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
