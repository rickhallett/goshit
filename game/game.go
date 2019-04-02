package game

import (
	"fmt"
	c "goshit/card"
	p "goshit/player"
)

type State struct {
	Players []*p.Player
	Active *p.Player
	Pile c.Pile
	Deck c.Deck
}

func (s *State) InitPlayers(n int, deck c.Deck) {
	for i := 0; i < n; i++ {
		p := &p.Player{
			Name: fmt.Sprintf("Player %v", i + 1),
		}

		var err *c.DeckError

		p.Hand, err = deck.Deal(3)
		if err != nil {
			err.AddMsg(fmt.Sprintf("Player name: %s", p.Name))
			err.AddMsg(fmt.Sprint("Attempted deal to hand"))
			fmt.Println(err)
		}

		p.Table, err = deck.Deal(3)
		if err != nil {
			err.AddMsg(fmt.Sprintf("Player name: %s", p.Name))
			err.AddMsg(fmt.Sprint("Attempted deal to table"))
			fmt.Println(err)
		}

		p.Blind, err = deck.Deal(3)
		if err != nil {
			err.AddMsg(fmt.Sprintf("Player name: %s", p.Name))
			err.AddMsg(fmt.Sprint("Attempted deal to blind"))
			fmt.Println(err)
		}

		if p.CardsReady() {
			fmt.Printf("%s is ready to play!\n", p.Name)
		}


		s.Players = append(s.Players, p)
	}
}

func InitState() *State {
	return &State{}
}