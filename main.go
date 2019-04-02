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
	//var deck card.Deck
	state := game.InitState()
	rawDeck := card.CreateRawDeck()
	state.Deck.Cards = card.Shuffle(rawDeck)
	state.InitPlayers(3, state.Deck)
	//util.PrettyPrint(state)

	p1 := state.Players[0]
	util.PrintPlayerHand(p1)
	//util.PrintPlayerCards(p1)
	//p1.SwitchCard(1, 1)
	//util.PrintPlayerCards(p1)
	p1.PlayCard(1, state)

}
