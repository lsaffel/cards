package main

import "fmt"

func main(){
	// var card string = "Ace of Spades"
	// this line is the equivalent to that line and this shorter version
	// is more often used to declare a variable: 
	// card := "Ace of Spades"

	card := newCard()

	fmt.Println(card)
}

func newCard() string{		// string here says what the function returns, a string
	return "Five of Diamonds"
}