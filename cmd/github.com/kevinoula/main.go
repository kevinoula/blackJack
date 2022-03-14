package main

import (
	"fmt"
	"cards"
	"player"
	// "os"
	// "bufio"
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
	

	// read user input
	fmt.Println("Press [enter] to start...")
	var input string
	fmt.Scanf("%s\n", &input)
	// fmt.Printf("Collected input: %v of type %T with length %v.\n", input, input, len(input))

	// game status
	playing := false
	if input == "" {
		playing = true
	}

	for playing {
		playRound()
		fmt.Println("Press [enter] to continue, anything else to stop.")
		fmt.Scanf("%s\n", &input)
		
		// end game when user wants to stop
		if input != "" {
			playing = false
		}
	}
}
