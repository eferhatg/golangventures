package main

func main() {

	cards := newDeckFromFile("initial_cards")
	cards.shuffle()
	cards.print()
	//	cards.saveToFile("initial_cards")
	// hand, remainingcards := deal(cards, 4)
	// hand.print()
	// remainingcards.print()

}
