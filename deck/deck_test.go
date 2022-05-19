package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but received %v", len(d))
	}
	if d[0] != "Ace of Spades" || d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected first/last card not matchin")
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	d.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected deck length of 16 but received %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
