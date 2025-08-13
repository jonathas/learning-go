package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, but got %s", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected last card to be King of Spades, but got %s", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	os.Remove(filename)

	deck := newDeck()
	deck.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}

	if loadedDeck[0] != "Ace of Hearts" {
		t.Errorf("Expected first card to be Ace of Hearts, but got %s", loadedDeck[0])
	}

	if loadedDeck[len(loadedDeck)-1] != "King of Spades" {
		t.Errorf("Expected last card to be King of Spades, but got %s", loadedDeck[len(loadedDeck)-1])
	}

	os.Remove(filename)
}