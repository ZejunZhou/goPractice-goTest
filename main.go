package main

import (
	"fmt"
)


func main(){
	fmt.Println("start of project")
    // cards := newDeck()
	// cards.print()
	// hand, remainingCards := deal(cards, 5)
	// fmt.Println("start of hands")
	// hand.print()
	// fmt.Println("start of remaining cards")
	// remainingCards.print()
	cards := newDeck()
	// fmt.Println(cards.toString())
	// cards.saveToFile("my_cards")
	// newDeckFromFile("my_cards").print()
	cards.shuffle()
	cards.print()
	
}

func newCard() string{
	return "ace"
}