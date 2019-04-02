package player

import (
	"fmt"
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

func (p *Player) SwitchCard(fromHand, toTable int) {
	handCard := p.Hand[fromHand - 1]
	tableCard := p.Table[toTable - 1]
	p.Hand[fromHand - 1] = tableCard
	p.Table[toTable - 1] = handCard
	fmt.Printf("%s swapped '%v %v' from their hand with the '%v %v' on the table\n",
		p.Name, handCard.Rank, handCard.Suit, tableCard.Rank, tableCard.Suit)
}

func (p *Player) PlayCard(num int, pile c.Pile) {

}