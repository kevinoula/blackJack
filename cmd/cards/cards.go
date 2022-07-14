package cards

import (
	"math/rand"
	"time"
)

// Card a basic blackjack card that has a name and suit.
type Card struct {
	Name  string
	Suit  string
	Value int
}

var valueMap = map[string]int{
	"Ace":   11,
	"2":     2,
	"3":     3,
	"4":     4,
	"5":     5,
	"6":     6,
	"7":     7,
	"8":     8,
	"9":     9,
	"10":    10,
	"Jack":  10,
	"Queen": 10,
	"King":  10,
}

var allowedSuits = []string{
	"spades",
	"hearts",
	"clubs",
	"diamonds",
}

type Deck struct {
	cards []Card
}

// NewDeck Initializes a new deck structure which contains a list containing 52 unique cards by default.
// Deck of cards can be accesses by Deck.cards as a slice.
func NewDeck() Deck {
	// Create a deck of each card for each suit
	var cards []Card
	for _, suit := range allowedSuits {
		for name := range valueMap {
			cards = append(cards, Card{Name: name, Suit: suit, Value: valueMap[name]})
		}
	}

	return Deck{cards}
}

// GetDeckLength Returns the number of cards in the deck.
func (d Deck) GetDeckLength() int {
	return len(d.cards)
}

// RemoveRandomCard Randomly removes a card from the deck.
func (d *Deck) RemoveRandomCard() Card {
	rand.Seed(time.Now().UnixNano())
	randomCardIdx := rand.Intn(len(d.cards))
	randomCard := d.cards[randomCardIdx]

	// remove card from deck
	d.cards = append(d.cards[:randomCardIdx], d.cards[randomCardIdx+1:]...)
	return randomCard
}
