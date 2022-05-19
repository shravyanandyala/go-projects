package main

func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
	newcards := newDeckFromFile("my_cards")
	newcards.shuffle()
	newcards.print()
}
