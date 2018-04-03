package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 16 {
		t.Errorf("Expected length of the deck is 16 but recieved %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("First element should be Ace of Spades but got %v", deck[0])
	}

	if deck[len(deck)-1] != "Four of Clubs" {
		t.Errorf("Last element should be Four of Clubs but got %v", deck[len(deck)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_savedDeck")
	deck := newDeck()
	deck.saveToFile("_savedDeck")
	decknew := newDeckFromFile("_savedDeck")
	if len(decknew) != 16 {
		t.Errorf("Expected 16 cards but got %v", len(decknew))
	}
	os.Remove("_savedDeck")
}
