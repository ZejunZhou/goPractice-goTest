package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
)

// Create a new type of 'deck', which is a slice of strings
//说明我们创了一个type deck, 作用和slice string 一样
type deck []string

func newDeck() deck{
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
    cards := deck{}
	for _, value := range cardValues{
		for _, suit := range cardSuits{
			cards = append(cards, value + " of " + suit)
		}
	}
	return cards
}

// receiver for print out card info
func (d deck) print() {
	for i, card := range(d){
		fmt.Println(i, card)
	}
}

// function split cards array into two seperate array
func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	cards := []string(d)
	cardsString := strings.Join(cards, ",")
	return cardsString
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if (err != nil) {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}


func (d deck) shuffle() {
	for i, _ := range(d){
		newPosition := rand.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}



