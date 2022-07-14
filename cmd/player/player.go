package player

import (
	"github.com/kevinoula/blackjack/cmd/cards"
)

type Player struct {
	Name            string
	Score           int
	HandValue       int
	AcesTransformed int
	hand            []cards.Card
}

// NewPlayer Initializes a new Player with an empty hand.
func NewPlayer(name string) Player {
	var newHand []cards.Card
	return Player{Name: name, Score: 0, HandValue: 0, AcesTransformed: 0, hand: newHand}
}

// AddToHand Adds a cards.Card to the hand.
func (p *Player) AddToHand(card cards.Card) {
	p.hand = append(p.hand, card)
}

// GetHand Returns a Player's entire hand.
func (p *Player) GetHand() []cards.Card {
	return p.hand
}

// EmptyHand Empties a player's hand.
func (p *Player) EmptyHand() {
	var newHand []cards.Card
	p.hand = newHand
}

// CountInHand Iteratively searches a Player's hand and counts the occurrences of a given card.
func (p Player) CountInHand(name string) int {
	/*
		Example:
		Input -> User.CountInHand(Ace)
		User's hand contains -> [{3 spades} {6 diamonds} {Ace clubs}]
		Output -> 1
	*/
	count := 0
	for _, card := range p.hand {
		cardName, _, _ := cards.GetCard(card)
		if name == cardName {
			count++
		}
	}
	return count
}
