package main

import (
	"os"
	"testing"
)

// t here is the test handler
func TestNewDeck (t *testing.T) {
	d := newDeck()

	// we have 4 suits and 4 values, so there should be 16 total cards
	if len(d) != 16 {
		// this  works. Apparently, the f represents formatting:
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}			

	if d[0] != "Ace of Spades" {
		// %v injects the value into the string
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0] )
	}
		
	if d[len(d) - 1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")	// removes any files we created 
								// in the past while testing

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}