package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 52 {
		t.Errorf("Expected 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades but got %s", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {

		t.Errorf("Expected King of Clubs but got %s", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_testingdeck")

	cards := newDeck()

	cards.saveToFile("_testingdeck")
	loadedDeck := newDeckFromFile("_testingdeck")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 but got %v", len(loadedDeck))
	}
	os.Remove("_testingdeck")
}
