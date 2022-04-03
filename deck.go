package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type, called 'deck', which is a slice of strings
// Now, any other file in this same package can use the type deck
type deck []string

func newDeck() deck {
	cards := deck{}		// creates a new variable of type cards of type deck 
						// and it starts off empty

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
		}
	}

	return cards
}

// print every card in the deck, plus i to show the index of what's printed
func (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

// takes a deck and converts it into a string separated by commas
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	// 0666 is the permissions. It means that anyone can read and write this file
}

func newDeckFromFile(filename string) deck {
	// bs is short for byte slice
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck() so that it returns a deck
		// Option #2 - log the error and entirely quit the program
		// We're going with option #2

		fmt.Println("Error:", err)
		os.Exit(1)			// a non-zero here indicates that there was an error
	}

	// string(bs)		// converts the byte slice into a string
					// becomes: Ace of Spades,Two of Spades,Three of Spades...

	s := strings.Split(string(bs), ",")
	return deck(s)		// deck(s)  turns the slice of strings into a deck

}

func (d deck) shuffle() {
	// create a random number generator that is truly random.
	// source is the seed that we pass to rand.New to specify the seed,
	// to make it truly random
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		// newPosition := rand.Intn(len(d) - 1)
		newPosition := r.Intn(len(d) - 1)

		// this swaps the elements at these two positions
		d[i], d[newPosition] = d[newPosition], d[i]
	}

}
