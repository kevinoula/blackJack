package player

import (
	"cards"
)

type Player struct {
	name string
	hand []cards.Card
	score int
	handValue int
}

func NewPlayer(name string) Player {
	/*
	Public Function.
	Initializes a new player with an empty hand.
	*/
	var newHand []cards.Card
	return Player{name, newHand, 0, 0}
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

func GetScore(player Player) int {
	/*
	Public Function.
	Returns a player's score.
	*/
	return player.score
}

func SetScore(player *Player, score int) {
	/*
	Public Function.
	Sets a player's score.
	*/
	player.score = score
}

func GetHandValue(player Player) int {
	/*
	Public Function.
	Returns a player's current hand value.
	*/
	return player.handValue
}

func SetHandValue(player *Player, value int) {
	/*
	Public Function.
	Sets a player's current hand value.
	*/
	player.handValue = value
}