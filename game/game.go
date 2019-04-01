package game

import (
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
			Name: string(i),
			Hand: deck.Deal(3),
			Table: deck.Deal(3),
			Blind: deck.Deal(3),
		}
		s.Players = append(s.Players, p)
	}
}

func InitState() *State {
	return &State{}
}