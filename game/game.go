package game

import (
	"fmt"
	c "goshit/card"
	p "goshit/player"
)

type State struct {
	Players []*p.Player
	Active *p.Player
	Pile []c.Card
}

func (s *State) InitPlayers(n int, deck c.Deck) {
	for i := 0; i < n; i++ {
		p := &p.Player{
			Name: fmt.Sprintf("Player %v", i + 1),
		}

		var err error

		p.Hand, err = deck.Deal(3)
		if err != nil {
			fmt.Println(err)
		}

		p.Table, err = deck.Deal(3)
		if err != nil {
			fmt.Println(err)
		}

		p.Blind, err = deck.Deal(3)
		if err != nil {
			fmt.Println(err)
		}


		s.Players = append(s.Players, p)
	}
}

func InitState() *State {
	return &State{}
}