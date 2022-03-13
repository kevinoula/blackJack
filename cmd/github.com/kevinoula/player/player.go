package player

import (
	"cards"
)

type Player struct {
	name string
	hand []cards.Card
}

func NewPlayer(name string) Player {
	/*
	Public Function.
	Initializes a new player with an empty hand.
	*/
	var newHand []cards.Card
	return Player{name, newHand}
}

func AddToHand(player *Player, card cards.Card) {
	/*
	Public Function.
	Adds a card to the player.hand slice.
	*/
	player.hand = append(player.hand, card)
}

func GetHand(player Player) []cards.Card {
	/*
	Public Function.
	Returns a player's entire player.hand slice.
	*/
	return player.hand
}

func GetName(player Player) string {
	/*
	Public Function.
	Returns a player's name.
	*/
	return player.name
}
