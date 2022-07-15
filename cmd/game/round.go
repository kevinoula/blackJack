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
	// otherwise subtract 10 so for each Ace so it esssentially updates it's value to 1
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
	fmt.Printf("[SYSTEM] There are %v cards in this deck\n", deck.GetDeckLength())
	fmt.Printf("--- Round %v ---\n", round)

	// Deal phase
	fmt.Println("--- Deal Phase ---")
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	deck, dealer = drawCard(deck, dealer)
	deck, user = drawCard(deck, user)
	fmt.Printf("[SYSTEM] %v's starting hand %v has a value of %v.\n", user.Name, user.GetHand(), user.HandValue)
	fmt.Printf("[SYSTEM] %v's starting hand %v has a value of %v.\n", dealer.Name, dealer.GetHand(), dealer.HandValue)
	fmt.Print("--- Deal Phase Ended ---\n\n")
	time.Sleep(5 * time.Second)

	// user's turn
	deck, user = playTurn(deck, user)

	// dealer's turn
	deck, dealer = playTurn(deck, dealer)

	// increment score
	fmt.Printf("[SYSTEM] %v's hand value is %v. %v's hand value is %v.\n", dealer.Name, dealer.HandValue, user.Name, user.HandValue)
	if user.HandValue > 21 && dealer.HandValue > 21 {
		// TIE if both over 21
		fmt.Println("[TIE] TIE since both players bust!")

	} else if user.HandValue > 21 {
		// DEALER WINS if user over 21
		fmt.Printf("[LOSE] %v has bust!\n", user.Name)
		dealer.Score = dealer.Score + 1

	} else if dealer.HandValue > 21 {
		// USER WINS if dealer over 21
		fmt.Printf("[WIN] %v has bust!\n", dealer.Name)
		user.Score = user.Score + 1

	} else if dealer.HandValue == user.HandValue {
		// Since nobody is over 21
		// TIE is both hands have same value
		fmt.Println("[TIE] TIE since both players hands are the same value!")

	} else if dealer.HandValue > user.HandValue {
		// DEALER WINS if hand value > than user's
		fmt.Printf("[LOSE] %v has the more valuable hand!\n", dealer.Name)
		dealer.Score = dealer.Score + 1

	} else {
		// USER WINS if hand value > than dealer's
		fmt.Printf("[WIN] %v has the more valuable hand!\n", user.Name)
		user.Score = user.Score + 1

	}

	// Output the current score
	user = resetPlayer(user)
	dealer = resetPlayer(dealer)
	round++
	fmt.Printf("[SCORE] %v has %v win(s). %v has %v win(s).\n", dealer.Name, dealer.Score, user.Name, user.Score)
	return
}
