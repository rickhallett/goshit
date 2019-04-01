package main

import (
	"goshit/card"
	"goshit/game"
	"goshit/util"
)

func init() {}

func main() {
	var deck card.Deck
	rawDeck := card.CreateRawDeck()
	deck.Cards = card.Shuffle(rawDeck)

	state := game.InitState()
	state.InitPlayers(10, deck)


	util.PrettyPrint(state)

}
