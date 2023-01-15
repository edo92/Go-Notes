package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func TestFileIO(t *testing.T) {
	const tempName = "_decktesting"
	os.Remove(tempName)

	deck := newDeck()
	deck.saveToFile((tempName))

	loadedDeck := newDeckFromFile(tempName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(tempName)
}
