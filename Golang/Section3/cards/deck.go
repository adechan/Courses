package main

import (
	"time"
	"math/rand"
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

// to compile + execute: go run main.go deck.go

// create a new type of 'deck'
// which is a slice of strings
type deck []string

// ^ it has all the behavior of []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string {"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string {"Ace", "Two", "Three", "Four"}

	for _, cardSuit := range cardSuits {
		cardItem := "";
		for _, cardValue := range cardValues {
			cardItem = cardValue + " of " + cardSuit
			cards = append(cards, cardItem)
		}
	}

	return cards
}

// RECEIVER function
func (d deck) print() {
	// d: convention (d is the first letter from deck)
	for _, card := range d {
		fmt.Println(card);
	}
}
// ^ any variable of type deck gets access to the print function
// d = the actual copy of the deck we are working with is available in the function as a variable called d
// deck = every variable of type deck can call this function on itself

// MULTIPLE RETURN VALUES
func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

// deck -> []string -> string -> []byte

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)

	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program

		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	s := strings.Split(string(byteSlice), ",")

	return deck(s)
}

func (d deck) shuffle() {
	t := time.Now()
	source := rand.NewSource(t.UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := 	r.Intn(len(d) - 1)
		
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}