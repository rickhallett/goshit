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