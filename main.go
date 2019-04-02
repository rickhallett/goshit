package main

import (
	"fmt"
	"goshit/game"
	"runtime"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	//var deck card.Deck
	state := game.InitState()
	rawDeck := game.CreateRawDeck()
	state.Deck.Cards = game.Shuffle(rawDeck)
	state.InitPlayers(3, &state.Deck)
	//util.PrettyPrint(state)

	p1 := state.Players[0]
	p2 := state.Players[1]
	p3 := state.Players[2]
	//game.PrintPlayerHand(p1)
	//util.PrintPlayerCards(p1)
	//p1.SwitchCard(1, 1)
	//util.PrintPlayerCards(p1)
	fmt.Println("Deck length:", len(state.Deck.Cards))
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)

	fmt.Println("Deck length:", len(state.Deck.Cards))
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	p1.PlayCard(1, state)
	p2.PlayCard(1, state)
	p3.PlayCard(1, state)
	game.PrintPlayerHand(p1)

}
