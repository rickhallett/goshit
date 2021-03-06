package game

import (
	"fmt"
)

type playerError struct {
	additionalMsgs []string
	err            error
}

func (pe *playerError) Error() string {
	return fmt.Sprintf("Player Error: %v %v", pe.additionalMsgs, pe.err)
}

func (pe *playerError) addMsg(msg string) {
	pe.additionalMsgs = append(pe.additionalMsgs, msg)
}

type Player struct {
	Name  string
	Hand  []Card
	Table []Card
	Blind []Card
	Turns int
}

func (p *Player) CardsReady() bool {
	if p.Hand != nil && p.Table != nil && p.Blind != nil {
		return true
	}

	return false
}

func (p *Player) SwitchCard(fromhand, totable int) {
	handCard := p.Hand[fromhand-1]
	tableCard := p.Table[totable-1]
	p.Hand[fromhand-1] = tableCard
	p.Table[totable-1] = handCard
	fmt.Printf("%s swapped '%v %v' from their hand with the '%v %v' on the table\n",
		p.Name, handCard.Rank, handCard.Suit, tableCard.Rank, tableCard.Suit)
}

// TODO: refactor this so that it has a common error pathway that returns to callee, less if/else statements
func (p *Player) PlayCard(num int, state *state) {
	var err *deckError // 2h, jd, 9s
	var newCard []Card

	playCard := p.Hand[num-1]
	if playCard == (Card{}) {
		err := fmt.Errorf("No card to play. Attempted to play card position %v", num)
		pe := &playerError{[]string{}, err}
		fmt.Println(pe)
	} else {
		p.Hand[num-1] = Card{}
		state.Pile.takeCard(playCard, p)
	}

	newCard, err = state.Deck.deal(1)
	if err != nil {
		err.addMsg(fmt.Sprintf("player name: %s", p.Name))
		err.addMsg(fmt.Sprint("Could not add another card to player hand"))
		fmt.Println(err)
	}

	if newCard != nil {
		p.Hand[num-1] = newCard[0]
		fmt.Printf("Added %v %v to %s's hand\n", newCard[0].Rank, newCard[0].Suit, p.Name)
	}

}
