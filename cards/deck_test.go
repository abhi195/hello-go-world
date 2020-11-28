package main

import (
	"os"
	"testing"
)

func TestCreateDeck(t *testing.T) {
	d := createDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Hearts" {
		t.Errorf("Expected first card of King of Hearts, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndLoadFromFile(t *testing.T) {
	// pre-cleanup
	os.Remove("_decktesting")

	// create & save
	deck := createDeck()
	deck.saveToFile("_decktesting")

	// load
	loadedDeck := loadFromFile("_decktesting")

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected deck length of %v, but got %v", len(deck), len(loadedDeck))
	}

	// post-cleanup
	os.Remove("_decktesting")
}
