package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Two of Spades" {
		t.Errorf("Expected first card of Two of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Ace of Clubs" {
		t.Errorf("Expected last card of Ace of Clubs, but got %v", d[len(d)-1])
	}
}

func TestDeal(t *testing.T) {
	d := newDeck()
	hand, remainingCards := deal(d, 10)

	if len(hand) != 10 {
		t.Errorf("Expected hand length of 10, but got %v", len(hand))
	}

	if len(remainingCards) != 42 {
		t.Errorf("Expected remaning card length of 42, but got %v", len(remainingCards))
	}
}

func TestSaveToDeckandNewDeckFromFile(t *testing.T) {
	// Cleanup old test file if exists
	os.Remove("_decktesting.tmp")

	d := newDeck()
	d.saveToFile("_decktesting.tmp")
	ld := newDeckFromFile("_decktesting.tmp")

	if len(ld) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(ld))
	}

	// Cleanup test files
	os.Remove("_decktesting.tmp")
}

func TestShuffle(t *testing.T) {
	d := newDeck()
	d.shuffle()

	if d[0] == "Two of Spades" && d[len(d)-1] == "Ace of Clubs" {
		t.Errorf("Expected cards to be shuffled")
	}
}
