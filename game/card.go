package game

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

// How to initialise: var err *c.DeckError, if err != nil {}
type deckError struct {
	additionalMsgs []string
	err            error
}

func (de *deckError) Error() string {
	return fmt.Sprintf("Deck Error: %v %v", de.additionalMsgs, de.err)
}

func (de *deckError) addMsg(msg string) {
	de.additionalMsgs = append(de.additionalMsgs, msg)
}

func (d *Deck) deal(n int) ([]Card, *deckError) {
	if len(d.Cards) < n {
		err := fmt.Errorf("Insufficient cards. Attempted to deal %v cards", n)
		de := &deckError{[]string{}, err}
		return nil, de
	}

	deal := d.Cards[len(d.Cards)-n:]
	d.Cards = d.Cards[:len(d.Cards)-n]
	return deal, nil
}

func (p *Pile) takeCard(c Card, pl *Player) {
	p.Cards = append(p.Cards, c)
	fmt.Printf("%s placed %v %v onto pile\n", pl.Name, c.Rank, c.Suit)
}

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
