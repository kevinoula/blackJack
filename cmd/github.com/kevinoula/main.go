package main

import (
	"fmt"
	"cards"
	"player"
)

func drawCard(currDeck cards.Deck, currPlayer player.Player) (cards.Deck, player.Player) {
	/*
	Private Function.
	Core function that draws a card from the deck and adds it to the player's hand.
	*/
	removedCard := cards.RemoveRandomCard(&currDeck)  // references the actual struct
	player.AddToHand(&currPlayer, removedCard)  // references the actual struct
	drawnValue, drawnSuit := cards.GetCard(removedCard)
	fmt.Printf("%v drew %v of %v\n", player.GetName(currPlayer), drawnValue, drawnSuit)
	fmt.Printf("%v's hand: %v\n", player.GetName(currPlayer), player.GetHand(currPlayer))
	return currDeck, currPlayer
}

func playRound()  {
	deck := cards.NewDeck()
	cards.GetDeck(deck)
	user, dealer := player.NewPlayer("User"), player.NewPlayer("Dealer")

	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	fmt.Printf("Player's starting hand: %v\n", player.GetHand(user))
	cards.GetDeck(deck)	
}

func main() {
	fmt.Println("Welcome to Blackjack Simulator feat. Go!")
	fmt.Println("Press [1] to start...")

	var input string
	fmt.Scanf("%s", &input)
	for input == "1" {
		playRound()
		fmt.Println("Press [1] to continue...")
		fmt.Scanf("%s", &input)
	}
}
