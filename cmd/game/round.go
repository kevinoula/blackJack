package game

import (
	"fmt"
	"github.com/kevinoula/blackjack/cmd/cards"
	"github.com/kevinoula/blackjack/cmd/player"
	"time"
)

// drawCard Utility function that draws a card from the deck and adds it to the player's hand.
func drawCard(currDeck cards.Deck, currPlayer *player.Player) (cards.Deck, *player.Player) {
	drawnCard := currDeck.RemoveRandomCard() // references the actual struct
	currPlayer.AddToHand(drawnCard)          // references the actual struct
	fmt.Printf("[SYSTEM] %v drew %v of %v\n", currPlayer.Name, drawnCard.Name, drawnCard.Suit)
	// Wow this was kind of interesting to solve:
	// Since an Ace can be transformed into a 1 or a 11, count the number of Aces - the number of transformations
	// done already so no duplicate transformations can be performed
	// if the drawn cards does not bust the hand then ignore
	// otherwise subtract 10 so for each Ace, so it essentially updates its value to 1
	// until the hand is no longer bust
	numAces := currPlayer.CountInHand("Ace")
	for count := 0; count < numAces-currPlayer.AcesTransformed; count++ {
		if currPlayer.HandValue+drawnCard.Value <= 21 {
			break
		}
		currPlayer.HandValue = currPlayer.HandValue - 10
		currPlayer.AcesTransformed = currPlayer.AcesTransformed + 1
	}
	currPlayer.HandValue = currPlayer.HandValue + drawnCard.Value
	time.Sleep(2 * time.Second)
	return currDeck, currPlayer
}

// resetPlayer Utility function that empties a player's hand and sets the value of the hand to 0 to start a new round.
func resetPlayer(currPlayer *player.Player) *player.Player {
	currPlayer.EmptyHand()
	currPlayer.HandValue = 0
	currPlayer.AcesTransformed = 0
	return currPlayer
}

// playTurn Utility function that orchestrates a player's turn.
// A round continues until a player's hand goes over 21 or they end their turn.
// A human player has the option to decide when to end their turn.
// The Dealer is programmed to call until their hand hits at least 18.
func playTurn(currDeck cards.Deck, currPlayer *player.Player) (cards.Deck, *player.Player) {
	fmt.Printf("--- It is %v's turn! ---\n", currPlayer.Name)
	fmt.Printf("[SYSTEM] %v's hand %v has a value of %v.\n", currPlayer.Name, currPlayer.GetHand(), currPlayer.HandValue)
	for currPlayer.HandValue <= 21 {
		// end turn if dealer hits limit
		if currPlayer.Name == "Dealer" && currPlayer.HandValue >= 18 {
			break

		} else if currPlayer.Name != "Dealer" {
			// end turn if user folds
			fmt.Println("[SYSTEM] Press [1] to draw card, [2] to end turn.")
			var input string
			_, err := fmt.Scanf("%s\n", &input)
			if err != nil {
				fmt.Printf("[SYSTEM] error reading input: %v\n", err)
				fmt.Println("[SYSTEM] Continuing as normal")
			}

			if input == "2" {
				break
			}
		}

		currDeck, currPlayer = drawCard(currDeck, currPlayer)
		fmt.Printf("[SYSTEM] %v's hand %v has a value of %v.\n", currPlayer.Name, currPlayer.GetHand(), currPlayer.HandValue)
		fmt.Printf("[SYSTEM] There are %v cards in this deck\n", currDeck.GetDeckLength())
	}
	fmt.Printf("--- %v's turn has ended ---\n\n", currPlayer.Name)
	time.Sleep(5 * time.Second)
	return currDeck, currPlayer
}

// PlayRound Core function that logically orchestrates a round of Blackjack.
// Updates the winning player's score and returns an incremented round number.
func PlayRound(user *player.Player, dealer *player.Player, round int) {
	deck := cards.NewDeck()

	// Deal phase
	fmt.Println("--- Deal Phase ---")
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	fmt.Printf("[SYSTEM] %v's starting hand %v has a value of %v.\n", user.Name, user.GetHand(), user.HandValue)
	fmt.Printf("[SYSTEM] %v's starting hand %v has a value of %v.\n", dealer.Name, dealer.GetHand(), dealer.HandValue)
	fmt.Printf("[SYSTEM] There are %v cards in this deck\n", deck.GetDeckLength())
	fmt.Print("--- Deal Phase Ended ---\n\n")
	time.Sleep(5 * time.Second)

	// user's turn
	deck, user = playTurn(deck, user)

	// dealer's turn
	deck, dealer = playTurn(deck, dealer)

	// increment score
	fmt.Printf("[SYSTEM] %v's hand value is %v. %v's hand value is %v.\n", dealer.Name, dealer.HandValue, user.Name, user.HandValue)
	if user.HandValue > 21 && dealer.HandValue > 21 || dealer.HandValue == user.HandValue { // Tie scenario
		fmt.Printf("[ROUND %d] ends in a TIE!\n", round)

	} else if user.HandValue > 21 || (21 >= dealer.HandValue && dealer.HandValue > user.HandValue) { // Dealer win scenario
		fmt.Printf("[ROUND %d] Dealer wins!\n", round)
		dealer.Score = dealer.Score + 1

	} else { // Otherwise user win scenario
		fmt.Printf("[ROUND %d] %s wins!\n", round, user.Name)
		user.Score = user.Score + 1
	}

	// Output the current score
	user = resetPlayer(user)
	dealer = resetPlayer(dealer)
	round++
	fmt.Printf("[SCORE] %v has %v win(s). %v has %v win(s).\n", dealer.Name, dealer.Score, user.Name, user.Score)
	return
}
