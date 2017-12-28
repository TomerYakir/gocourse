package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string // like saying deck extends slice of string

func newDeck() deck { // doesn't need a receiver because it won't be an "instance-level" method
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ones", "Twos", "Threes", "Fours", "Fives", "Sixes", "Sevens", "Eights", "Nines",
		"Tens", "Jacks", "Queens", "Kings", "Aces"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, suit+" of "+value)
		}
	}
	return cards
}

func (d deck) toString() string {
	// deck -> []string -> string
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() deck {

	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		rnd := r.Intn(len(d) - 1)
		d[i], d[rnd] = d[rnd], d[i] //swap elements
	}
	return d
}

func readDeckFromFile(filename string) deck {
	fileData, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("error reading from file ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(fileData), ","))
}

func (d deck) deal(size int) (deck, deck) { // returns a hand and the sliced deck
	return d[:size], d[size:]
}

func (d deck) print() { // the reason we refer to the instance as "d" is for convention
	for i, card := range d {
		fmt.Println(i, card)
	}
}
