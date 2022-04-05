package main

import (
	"os"
	"testing"
)

// t here is the test handler
func TestNewDeck (t *testing.T) {
	d := newDeck()

	// we have 4 suits and 13 values, so there should be 52 total cards
	if len(d) != 52 {
		// this  works. The f represents formatting:
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}			

	// check the first card in the deck to make sure it's correct
	if d[0] != "Ace of Spades" {
		// %v injects the value into the string
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0] )
	}
	
	// check the last card in the deck
	if d[len(d) - 1] != "King of Clubs" {
		t.Errorf("Expected last card of King of Clubs, but got %v", d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")	// removes any files we created 
								// in the past while testing

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	// make sure the deck has the correct number of cards in it, 52
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")
}