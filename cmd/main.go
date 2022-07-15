package main

import (
	"fmt"
	"github.com/kevinoula/blackjack/cmd/game"
	"github.com/kevinoula/blackjack/cmd/player"
)

func main() {
	// init game
	fmt.Println("[SYSTEM] Welcome to Blackjack Simulator feat. Go!")
	fmt.Println("[SYSTEM] Enter player's name:")
	var input string
	_, err := fmt.Scanf("%s\n", &input)
	if err != nil {
		fmt.Printf("[SYSTEM] error reading input: %v\n", err)
		fmt.Println("[SYSTEM] Continuing as normal")
	}

	user, dealer := player.NewPlayer(input), player.NewPlayer("Dealer")
	playing := true
	round := 1

	for playing {
		game.PlayRound(&user, &dealer, round)
		round += 1
		fmt.Println("[SYSTEM] Press [1] to keep playing, anything else to stop.")
		_, err = fmt.Scanf("%s\n", &input)
		if err != nil {
			fmt.Printf("[SYSTEM] error reading input: %v\n", err)
			fmt.Println("[SYSTEM] Continuing as normal")
		}

		// end game when user wants to stop
		if input != "1" {
			playing = false
		}
	}
}
