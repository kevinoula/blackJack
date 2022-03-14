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
	drawnCard, drawnSuit, drawnValue := cards.GetCard(removedCard)
	fmt.Printf("[DRAW] %v drew %v of %v\n", player.GetName(currPlayer), drawnCard, drawnSuit)
	player.SetHandValue(&currPlayer, player.GetHandValue(currPlayer)+drawnValue)
	// fmt.Printf("%v's hand %v has value %v.\n", player.GetName(currPlayer), player.GetHand(currPlayer), player.GetHandValue(currPlayer))
	return currDeck, currPlayer
}

func playRound()  {
	deck := cards.NewDeck()
	cards.GetDeck(deck)
	user, dealer := player.NewPlayer("User"), player.NewPlayer("Dealer")

	// Deal phase 
	fmt.Println("--- Deal Phase ---")
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	fmt.Printf("Player's starting hand %v has a value of %v and score of %v.\n", player.GetHand(user), player.GetHandValue(user), player.GetScore(user))
	fmt.Printf("Dealer's starting hand %v has a value of %v and score of %v.\n", player.GetHand(dealer), player.GetHandValue(dealer), player.GetScore(dealer))
	fmt.Print("--- Deal Phase Ended ---\n\n")
	// cards.GetDeck(deck)	

	// round continues until a player's hand goes over 21 or they end their turn
	// player user's turn
	fmt.Printf("--- It is %v's turn! ---\n", player.GetName(user))
	for player.GetHandValue(user) <= 21 {
		fmt.Printf("%v's hand %v has a value of %v.\n", player.GetName(user), player.GetHand(user), player.GetHandValue(user))

		// end turn if user folds
		fmt.Println("Press [1] to draw card, [2] to end turn.")
		var input string
		fmt.Scanf("%s\n", &input)
		if input == "2" {
			break
		}
		deck, user = drawCard(deck, user)
		cards.GetDeck(deck)	
	}
	fmt.Printf("--- %v's turn has ended ---\n\n", player.GetName(user))

	// play dealer's turn
	fmt.Printf("--- It is %v's turn ---\n", player.GetName(dealer))
	for player.GetHandValue(dealer) <= 21 {
		fmt.Printf("%v's hand %v has a value of %v.\n", player.GetName(dealer), player.GetHand(dealer), player.GetHandValue(dealer))

		// end turn if dealer has a hand value of at least 18
		if  player.GetHandValue(dealer) >= 18 {
			break
		}
		deck, dealer = drawCard(deck, dealer)
		cards.GetDeck(deck)	
	}
	fmt.Printf("--- %v's turn has ended ---\n\n", player.GetName(dealer))

	// Check if any player is over 21
	// TIE if both over 21
	// USER WINS if dealer over 21
	// DEALER WINS if user over 21

	// Since nobody is over 21
	// TIE is both hands have same value
	// USER WINS if hand value > than dealer's
	// DEALER WINS if hand value > user's

	// increment score
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
