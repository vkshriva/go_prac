/*
Create a list of playing cards.  Essentially an array of strings
Log out the contents of a deck of cards
Shuffles all the cards in a deck
Create a 'hand' of cards.
Save a list of cards to a file on the local machine
Load a list of cards from the local machine
*/
package main

func main() {
	cards := newCard()

	handsCard, _ := deal(cards, 5)

	handsCard.saveToFile("handsCard")
	newHandsCard := newDeckFromFile("handsCard")

	newHandsCard.shuffleCard()
	newHandsCard.print()

}
