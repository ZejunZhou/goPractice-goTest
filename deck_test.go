package main

import (
	"testing"
	"os"
	)

func TestNewDeck(t *testing.T){
	d := newDeck()
	if (len(d) != 16){
		t.Errorf("Expect length of 16 for deck, but get %d", len(d))
	}
}

func TestSaveToDeck(t *testing.T){
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadDeck := newDeckFromFile("_decktesting")

	if len([]string(loadDeck)) != 16 {
		t.Errorf("Expected length should be 16, but get %d", len([]string(loadDeck)))
	}
	os.Remove("_decktesting")
}