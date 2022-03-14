package cards

import (
	"fmt"
	"math/rand"
	"time"
)

type Card struct {
	value string
	suit  string
}

var valueMap = map[string]int {
	"Ace": 11,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"10": 10,
	"Jack": 10,
	"Queen": 10,
	"King": 10,
}

// var allowedValues = []string{
// 	"Ace",
// 	"2",
// 	"3",
// 	"4",
// 	"5",
// 	"6",
// 	"7",
// 	"8",
// 	"9",
// 	"10",
// 	"Jack",
// 	"Queen",
// 	"King",
// }

var allowedSuits = []string {
	"spades",
	"hearts",
	"clubs",
	"diamonds",
}


func newCard(value string, suit string) Card {
	/*
	Private function.
	Initializes a new card structure. A card has a value and a suit.
	*/
	return Card{value, suit}
}

func GetCard(card Card) (string, string, int) {
	/*
	Public Function.
	Returns the value and suit of a given card.
	*/

	return card.value, card.suit, valueMap[card.value]
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
			newDeck  = append(newDeck, newCard(value, suit))
		}
	}

	return Deck{newDeck}
}

func GetDeck(deck Deck)  {
	/*
	Public Function.
	Prints out each card in a deck and the total size of the deck.
	*/
	// for _, card := range deck.cards {
	// 	fmt.Printf("%v of %v\n", card.value, card.suit)
	// }
	fmt.Printf("There are %v cards in this deck\n", len(deck.cards))
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
