package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16 but got %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected 'Ace of Spades' but got %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected 'Four of Clubs' but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")
	d := newDeck()
	err := d.saveToFile("_decktesting")
	if err != nil {
		t.Errorf("Expected nil but got err %v", err)
	}
	ud := newDeckFromFile("_decktesting")
	if ud == nil {
		t.Errorf("Expected deck file to exist")
	}
	if len(ud) != 16 {
		t.Errorf("Expected deck length to be 16 but got %v", len(ud))
	}
	os.Remove("_decktesting")
}
