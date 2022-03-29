package main

import "fmt"

func main(){

	cards := newDeck()

	// the deal function is in the deck.go file
	hand, remainingCards := deal(cards, 5)	

	hand.print()
	fmt.Println("----------------")

	remainingCards.print()
}
