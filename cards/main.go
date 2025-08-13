package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remaining := deal(cards, 5)
	hand.print()
	remaining.print()
	fmt.Println(cards.toString())
	cards.saveToFile("my_cards.txt")
	
	newCards := newDeckFromFile("my_cards.txt")
	newCards.print()

	newCards.shuffle()
	newCards.print()
}
