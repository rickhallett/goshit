package main

import (
	"goshit/card"
	"goshit/game"
	"goshit/util"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	var deck card.Deck
	rawDeck := card.CreateRawDeck()
	deck.Cards = card.Shuffle(rawDeck)

	state := game.InitState()
	state.InitPlayers(3, deck)
	//util.PrettyPrint(state)

	p1 := state.Players[0]
	util.PrintPlayerHand(p1)
	//util.PrintPlayerCards(p1)
	//p1.SwitchCard(1, 1)
	//util.PrintPlayerCards(p1)

}
