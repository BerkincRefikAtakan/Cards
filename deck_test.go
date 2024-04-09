package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("expected cards number 16, but there is %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card is Ace of Spades, but you have %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card is Four of Clubs,but you have %v", d[len(d)-1])
	}
}

func testSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck ,but there is %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
