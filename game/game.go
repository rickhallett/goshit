package game

import (
	"fmt"
)

type state struct {
	Players []*Player
	Active  *Player
	Pile    Pile
	Deck    Deck
}

func (s *state) InitPlayers(n int, deck *Deck) {
	for i := 0; i < n; i++ {
		p := &Player{
			Name: fmt.Sprintf("Player %v", i+1),
		}

		var err *deckError

		p.Hand, err = deck.deal(3)
		if err != nil {
			err.addMsg(fmt.Sprintf("Player name: %s", p.Name))
			err.addMsg(fmt.Sprint("Attempted deal to hand"))
			fmt.Println(err)
		}

		p.Table, err = deck.deal(3)
		if err != nil {
			err.addMsg(fmt.Sprintf("Player name: %s", p.Name))
			err.addMsg(fmt.Sprint("Attempted deal to table"))
			fmt.Println(err)
		}

		p.Blind, err = deck.deal(3)
		if err != nil {
			err.addMsg(fmt.Sprintf("Player name: %s", p.Name))
			err.addMsg(fmt.Sprint("Attempted deal to blind"))
			fmt.Println(err)
		}

		if p.CardsReady() {
			fmt.Printf("%s is ready to play!\n", p.Name)
		}

		s.Players = append(s.Players, p)
	}
}

func InitState() *state {
	return &state{}
}
