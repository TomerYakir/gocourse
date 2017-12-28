package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) { // note that here the function starts with upper-case (to export)
	d := newDeck()
	if len(d) != 56 {
		t.Errorf("Expected deck length of 56, but got %v", len(d))
	}
}

func TestDeckContent(t *testing.T) {
	d := newDeck()
	first := "Spades of Ones"
	last := "Clubs of Aces"
	if d[0] != first {
		t.Errorf("Expected first card of %v, but got %v", first, d[0])
	}
	if d[len(d)-1] != last {
		t.Errorf("Expected last card of %v, but got %v", last, d[len(d)-1])
	}
}

func cleanupTestFiles(filename string) {
	os.Remove(filename)
}

func TestSaveToDeckAndReadFromFile(t *testing.T) {
	filename := "_decktesting"
	// cleanup first
	cleanupTestFiles(filename)
	// create, save, load, check length, cleanup
	d := newDeck()
	d.saveToFile(filename)
	newd := readDeckFromFile(filename)
	if len(d) != len(newd) {
		t.Errorf("Expected length of %v, but got %v", len(d), len(newd))
	}
	cleanupTestFiles(filename)
}
