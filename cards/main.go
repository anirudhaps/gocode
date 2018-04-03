package main

func main() {
	//cards := newDeck()
	//fmt.Println("Deck before deal:")
	//cards.print()
	//hand, remainingCards := deal(cards, 4)
	//fmt.Println("New hand of cards:")
	//hand.print()
	//fmt.Println("Remaining cards in the deck after deal:")
	//remainingCards.print()
	//cards := newDeck()
	//fmt.Println(cards.toString())
	//cards.saveToFile("my_cards")
	//cards := newDeckFromFile("my_cards")
	//cards.print()
	cards := newDeck()
	cards.shuffle()
	cards.print()
}
