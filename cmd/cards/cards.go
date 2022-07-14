package cards

import (
	"math/rand"
	"time"
)

type Card struct {
	name string
	suit string
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

func newCard(value string, suit string) Card {
	/*
		Private function.
		Initializes a new card structure. A card has a name and a suit.
	*/
	return Card{value, suit}
}

func GetCard(card Card) (string, string, int) {
	/*
		Public Function.
		Returns the name, suit, and value of a given card.
	*/

	return card.name, card.suit, valueMap[card.name]
}

type Deck struct {
	cards []Card
}

func NewDeck() Deck {
	/*
		Public Function.
		Initializes a new deck structure which contains a list containing 52 unique cards by default.
		Deck of cards can be accesses by Deck.cards as a slice.
	*/
	// Create a deck of each card for each suit
	var newDeck []Card
	for _, suit := range allowedSuits {
		for value := range valueMap {
			newDeck = append(newDeck, newCard(value, suit))
		}
	}

	return Deck{newDeck}
}

func GetDeck(deck Deck) int {
	/*
		Public Function.
		Returns the number of cards in the deck.
	*/
	// for _, card := range deck.cards {
	// 	fmt.Printf("%v of %v\n", card.name, card.suit)
	// }
	// fmt.Printf("There are %v cards in this deck\n", len(deck.cards))
	return len(deck.cards)
}

func RemoveRandomCard(deck *Deck) Card {
	/*
		Public Function.
		Randomly removes a card from the deck.
	*/
	rand.Seed(time.Now().UnixNano())
	randomCardIdx := rand.Intn(len(deck.cards))
	randomCard := deck.cards[randomCardIdx]

	// remove card from deck
	deck.cards = append(deck.cards[:randomCardIdx], deck.cards[randomCardIdx+1:]...)
	return randomCard
}
