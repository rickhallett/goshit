package player

import (
	"fmt"
	c "goshit/card"
	"goshit/game"
)

type Player struct {
	Name  string
	Hand  []c.Card
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
	handCard := p.Hand[fromHand-1]
	tableCard := p.Table[toTable-1]
	p.Hand[fromHand-1] = tableCard
	p.Table[toTable-1] = handCard
	fmt.Printf("%s swapped '%v %v' from their hand with the '%v %v' on the table\n",
		p.Name, handCard.Rank, handCard.Suit, tableCard.Rank, tableCard.Suit)
}

func (p *Player) PlayCard(num int, state *game.State) {
	var err *c.DeckError
	var newCard []c.Card

	playCard := p.Hand[num-1]
	state.Pile.TakeCard(playCard)

	newCard, err = state.Deck.Deal(1)
	if err != nil {
		err.AddMsg(fmt.Sprintf("Player name: %s", p.Name))
		err.AddMsg(fmt.Sprint("Could not add another card to player hand"))
		fmt.Println(err)
	}

	if newCard != nil {
		p.Hand[num-1] = newCard[0]
		fmt.Printf("Added %v %v to %s's hand", newCard[0].Rank, newCard[0].Suit, p.Name)
	}

}
