package main

import (
	"fmt"
	"cards"
	"player"
)

func drawCard(currDeck cards.Deck, currPlayer player.Player) (cards.Deck, player.Player) {
	/*
	Private Function.
	Utility function that draws a card from the deck and adds it to the player's hand.
	*/
	removedCard := cards.RemoveRandomCard(&currDeck)  // references the actual struct
	player.AddToHand(&currPlayer, removedCard)  // references the actual struct
	drawnCard, drawnSuit, drawnValue := cards.GetCard(removedCard)
	fmt.Printf("[SYSTEM] %v drew %v of %v\n", player.GetName(currPlayer), drawnCard, drawnSuit)
	// Wow this was kind of interesting to solve:
	// Since an Ace can be transformed into a 1 or a 11, count the number of Aces - the number of transformations
	// done already so no duplicate transformations can be performed
	// if the drawn cards does not bust the hand then ignore
	// otherwise subtract 10 so for each Ace so it esssentially updates it's value to 1 
	// until the hand is no longer bust
	numberofAces :=  player.CountInHand(currPlayer, "Ace")
	for count := 0; count < numberofAces - player.GetAcesTransformed(currPlayer); count++ {
		if player.GetHandValue(currPlayer)+drawnValue <= 21 {
			break
		}
		player.SetHandValue(&currPlayer, player.GetHandValue(currPlayer)-10)
		player.SetAcesTransformed(&currPlayer, player.GetAcesTransformed(currPlayer)+1)
	}

	player.SetHandValue(&currPlayer, player.GetHandValue(currPlayer)+drawnValue)
	return currDeck, currPlayer
}

func resetPlayer(currPlayer player.Player) player.Player {
	/*
	Private Function.
	Utility function that empties a player's hand and sets the value of the hand to 0 to start a new round.
	*/
	player.EmptyHand(&currPlayer)
	player.SetHandValue(&currPlayer, 0)
	player.SetAcesTransformed(&currPlayer, 0)
	return currPlayer
}

func playRound(user player.Player, dealer player.Player, round int) (player.Player, player.Player, int)  {
	/*
	Private Function.
	Core function that logically orchestrates a round of Blackjack.
	Updates the winning player's score and returns an incremented round number.
	*/
	deck := cards.NewDeck()
	fmt.Printf("[SYSTEM] There are %v cards in this deck\n", cards.GetDeck(deck))	
	fmt.Printf("--- Round %v ---\n", round)

	// Deal phase 
	fmt.Println("--- Deal Phase ---")
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	fmt.Printf("[SYSTEM] %v's starting hand %v has a value of %v and score of %v.\n", player.GetName(user), player.GetHand(user), player.GetHandValue(user), player.GetScore(user))
	fmt.Printf("[SYSTEM] %v's starting hand %v has a value of %v and score of %v.\n", player.GetName(dealer), player.GetHand(dealer), player.GetHandValue(dealer), player.GetScore(dealer))
	fmt.Print("--- Deal Phase Ended ---\n\n")

	// round continues until a player's hand goes over 21 or they end their turn
	// user's turn
	fmt.Printf("--- It is %v's turn! ---\n", player.GetName(user))
	fmt.Printf("[SYSTEM] %v's hand %v has a value of %v.\n", player.GetName(user), player.GetHand(user), player.GetHandValue(user))
	for player.GetHandValue(user) <= 21 {
		// fmt.Printf("[SYSTEM] %v's hand %v has a value of %v.\n", player.GetName(user), player.GetHand(user), player.GetHandValue(user))

		// end turn if user folds
		fmt.Println("[SYSTEM] Press [1] to draw card, [2] to end turn.")
		var input string
		fmt.Scanf("%s\n", &input)
		if input == "2" {
			break
		}
		deck, user = drawCard(deck, user)
		fmt.Printf("[SYSTEM] %v's hand %v has a value of %v.\n", player.GetName(user), player.GetHand(user), player.GetHandValue(user))
		fmt.Printf("[SYSTEM] There are %v cards in this deck\n", cards.GetDeck(deck))
	}
	fmt.Printf("--- %v's turn has ended ---\n\n", player.GetName(user))

	// play dealer's turn
	fmt.Printf("--- It is %v's turn ---\n", player.GetName(dealer))
	fmt.Printf("%v's hand %v has a value of %v.\n", player.GetName(dealer), player.GetHand(dealer), player.GetHandValue(dealer))
	for player.GetHandValue(dealer) <= 21 {
		// fmt.Printf("%v's hand %v has a value of %v.\n", player.GetName(dealer), player.GetHand(dealer), player.GetHandValue(dealer))

		// end turn if dealer has a hand value of at least 18
		if  player.GetHandValue(dealer) >= 18 {
			break
		}
		deck, dealer = drawCard(deck, dealer)
		fmt.Printf("%v's hand %v has a value of %v.\n", player.GetName(dealer), player.GetHand(dealer), player.GetHandValue(dealer))
		fmt.Printf("[SYSTEM] There are %v cards in this deck\n", cards.GetDeck(deck))
	}
	fmt.Printf("--- %v's turn has ended ---\n\n", player.GetName(dealer))

	// increment score
	fmt.Printf("[SYSTEM] %v's hand value is %v. %v's hand value is %v.\n", player.GetName(dealer), player.GetHandValue(dealer), player.GetName(user), player.GetHandValue(user))
	if player.GetHandValue(user) > 21 && player.GetHandValue(dealer) > 21 {
		// TIE if both over 21
		fmt.Println("[TIE] TIE since both players bust!")

	} else if player.GetHandValue(user) > 21 {
		// DEALER WINS if user over 21
		fmt.Printf("[LOSE] %v has bust!\n", player.GetName(user))
		player.SetScore(&dealer, player.GetScore(dealer)+1) 

	} else if player.GetHandValue(dealer) > 21 {
		// USER WINS if dealer over 21
		fmt.Printf("[WIN] %v has bust!\n", player.GetName(dealer))
		player.SetScore(&user, player.GetScore(user)+1) 

	} else if player.GetHandValue(dealer) == player.GetHandValue(user) {
		// Since nobody is over 21
		// TIE is both hands have same value
		fmt.Println("[TIE] TIE since both players hands are the same value!")

	} else if player.GetHandValue(dealer) > player.GetHandValue(user) {
		// DEALER WINS if hand value > than user's
		fmt.Printf("[LOSE] %v has the more valuable hand!\n", player.GetName(dealer))
		player.SetScore(&dealer, player.GetScore(dealer)+1) 

	} else {
		// USER WINS if hand value > than dealer's
		fmt.Printf("[WIN] %v has the more valuable hand!\n", player.GetName(user))
		player.SetScore(&user, player.GetScore(user)+1) 

	}

	// Output the current score
	user = resetPlayer(user)
	dealer = resetPlayer(dealer)
	round++
	fmt.Printf("[SCORE] %v has %v win(s). %v has %v win(s).\n", player.GetName(dealer), player.GetScore(dealer), player.GetName(user), player.GetScore(user))
	return user, dealer, round
}

func main() {
	// init game
	fmt.Println("[SYSTEM] Welcome to Blackjack Simulator feat. Go!")
	fmt.Println("[SYSTEM] Enter player's name:")
	var input string
	fmt.Scanf("%s\n", &input)
	user, dealer := player.NewPlayer(input), player.NewPlayer("Dealer")
	playing := true
	round := 1
	

	for playing {
		user, dealer, round= playRound(user, dealer, round)
		fmt.Println("[SYSTEM] Press [1] to keep playing, anything else to stop.")
		fmt.Scanf("%s\n", &input)
		
		// end game when user wants to stop
		if input != "1" {
			playing = false
		}
	}
}
