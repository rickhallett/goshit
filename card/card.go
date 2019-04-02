package card

import (
	"fmt"
	"math/rand"
	"time"
)

type Rank string
type Suit string

type Card struct {
	Rank Rank
	Suit Suit
}

type Deck struct {
	Cards []Card
}

type Pile struct {
	Cards []Card
}

type DeckError struct {
	desc, function string
	err error
}

func (de *DeckError) Error() string {
	return fmt.Sprintf("Deck Error: %v %v %v", de.desc, de.function, de.err)
}

// TODO: Deal is going to be used program wide in a number of places. Does it need to prevent errors itself (ie no cards left?)
func (d *Deck) Deal(n int) ([]Card, error) {
	if(len(d.Cards) < n) {
		err := fmt.Errorf("Insufficient cards. Attempted to deal %v cards", n)
		de := &DeckError{"DeckError", "func (d * Deck) Deal(n int) []Card, error {}", err}
		return nil, de
	}

	deal := d.Cards[len(d.Cards) - 3:]
	d.Cards = d.Cards[:len(d.Cards) - 3]
	return deal, nil
}

// Contains original deck of cards, unshuffled
var RawDeck []Card

// Shuffle creates a random permutation of index values from len(vals) and uses these to re-assign vals positions
func Shuffle(vals []Card) []Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	sDeck := make([]Card, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		sDeck[i] = vals[randIndex]
	}
	return sDeck
}

func CreateRawDeck() []Card {
	return []Card{
		{"2", "Spades"},
		{"3", "Spades"},
		{"4", "Spades"},
		{"5", "Spades"},
		{"6", "Spades"},
		{"7", "Spades"},
		{"8", "Spades"},
		{"9", "Spades"},
		{"10", "Spades"},
		{"J", "Spades"},
		{"Q", "Spades"},
		{"K", "Spades"},
		{"A", "Spades"},

		{"2", "Hearts"},
		{"3", "Hearts"},
		{"4", "Hearts"},
		{"5", "Hearts"},
		{"6", "Hearts"},
		{"7", "Hearts"},
		{"8", "Hearts"},
		{"9", "Hearts"},
		{"10", "Hearts"},
		{"J", "Hearts"},
		{"Q", "Hearts"},
		{"K", "Hearts"},
		{"A", "Hearts"},

		{"2", "Clubs"},
		{"3", "Clubs"},
		{"4", "Clubs"},
		{"5", "Clubs"},
		{"6", "Clubs"},
		{"7", "Clubs"},
		{"8", "Clubs"},
		{"9", "Clubs"},
		{"10", "Clubs"},
		{"J", "Clubs"},
		{"Q", "Clubs"},
		{"K", "Clubs"},
		{"A", "Clubs"},

		{"2", "Diamonds"},
		{"3", "Diamonds"},
		{"4", "Diamonds"},
		{"5", "Diamonds"},
		{"6", "Diamonds"},
		{"7", "Diamonds"},
		{"8", "Diamonds"},
		{"9", "Diamonds"},
		{"10", "Diamonds"},
		{"J", "Diamonds"},
		{"Q", "Diamonds"},
		{"K", "Diamonds"},
		{"A", "Diamonds"},
	}
}
