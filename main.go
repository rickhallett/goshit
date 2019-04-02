package main

import (
	"goshit/card"
	"goshit/game"
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
	state.InitPlayers(7, deck)


	//util.PrettyPrint(state)

}
