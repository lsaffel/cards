package main

func main(){

	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	cards.print()	// this function is in the deck.go file
}

func newCard() string{		// string here says what the function returns, a string
	return "Five of Diamonds"
}