package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 40 {
		t.Errorf("Expected length of 40 but got: %v", len(d))
	}

	if d[0] != "1 of Club" {
		t.Errorf("Expected first as 1 of Club: %v", d[0])
	}

	if d[len(d)-1] != "10 of Spade" {
		t.Errorf("Expected las as 10 of Spade: %v", d[len(d)-1])
	}
}

func TestSaveToFileAndGetDeckFromFile(t *testing.T) {
	testFileName := "_testingForDeckSaveLoadFromFile"
	os.Remove(testFileName)
	d := newDeck()
	d.shuffle()
	d.saveToFile(testFileName)
	loadedDeck := getDeckFromFile(testFileName)
	os.Remove(testFileName)
	if len(d) != len(loadedDeck) {
		t.Errorf("Loaded deck not equal to save deck")
	}
}
