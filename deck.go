package main

import "fmt"

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