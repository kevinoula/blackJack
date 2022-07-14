package player

import (
	"github.com/kevinoula/blackjack/cmd/cards"
)

type Player struct {
	name            string
	hand            []cards.Card
	score           int
	handValue       int
	acesTransformed int
}

func NewPlayer(name string) Player {
	/*
		Public Function.
		Initializes a new player with an empty hand.
	*/
	var newHand []cards.Card
	return Player{name, newHand, 0, 0, 0}
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

func EmptyHand(player *Player) {
	/*
		Public Function
		Empties a player's hand.
	*/
	var newHand []cards.Card
	player.hand = newHand
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

func CountInHand(player Player, cardValue string) int {
	/*
		Public Function.
		Iteratively searches a player's hand counts the occurances of a card value.

		Example:
		Input -> user, Ace
		User's hand contains -> [{3 spades} {6 diamonds} {Ace clubs}]
		Output -> 1
	*/
	count := 0
	for _, card := range player.hand {
		cardName, cardSuit, cardVal := cards.GetCard(card)
		_, _ = cardSuit, cardVal
		if cardValue == cardName {
			count++
		}
	}
	return count
}

func GetAcesTransformed(player Player) int {
	/*
		Public Function.
		Gets the number of Aces transformed.
	*/
	return player.acesTransformed
}

func SetAcesTransformed(player *Player, value int) {
	/*
		Public Function.
		Sets the number of Aces transformed.
	*/
	player.acesTransformed = value
}
