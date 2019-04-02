package player

import (
	c "goshit/card"
)

type Player struct {
	Name string
	Hand []c.Card
	Table []c.Card
	Blind []c.Card
	Turns int
}

func (p *Player) CardsReady() bool {
	if p.Hand != nil && p.Table != nil && p.Blind != nil {
		return true
	}

	return false
}