package main

import (
	"fmt"
)

/*
Create a new type of 'deck', which is a slice of
strings, that we can extend with receiver func's
*/
type deck []string

/*
Receiver; think method on any instance of deck
ex. card := deck{"card 1", "card 2"}
    card.print()
*/
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
	}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}
