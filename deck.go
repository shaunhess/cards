package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
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

/*
Receiver; converts deck to a single string
*/
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
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

// return two decks
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// Save deck to a file on disk
func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

// Load deck from file on disk
func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Log error and quit application
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

// Shuffle the deck of cards
func (d deck) shuffle() {
	// Need to generate a pseudo-rendom source for our seed
	source := rand.NewSource(time.Now().UnixNano())
	// Use our seed for random number generator
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
