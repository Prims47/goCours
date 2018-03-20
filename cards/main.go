package main

func main() {
	// cards := newDeck()

	// cards.saveToFile("my_cards")
	//fmt.Println(cards.toString())

	// cards := newDeckFromFile("my_cards")
	// cards.print()

	cards := newDeck()
	cards.shuffle()
	cards.print()
}
