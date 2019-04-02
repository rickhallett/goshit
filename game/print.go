package game

import (
	"encoding/json"
	"fmt"
)

func PrettyPrint(message ...interface{}) {
	for _, m := range message {
		jsonMessage, _ := json.MarshalIndent(m, "", "  ")
		fmt.Println(string(jsonMessage))
	}
}

func PrintPlayerCards(p *Player) {
	fmt.Printf("\nPlayer hand: \t1: %v %v \t2: %v %v \t3: %v %v \n"+
		"Player table: \t1: %v %v \t2: %v %v \t3: %v %v \n"+
		"Player Blind: \t1: ? ? \t\t2: ? ? \t\t3: ? ?\n\n",
		p.Hand[0].Rank, p.Hand[0].Suit,
		p.Hand[1].Rank, p.Hand[1].Suit,
		p.Hand[2].Rank, p.Hand[2].Suit,
		p.Table[0].Rank, p.Table[0].Suit,
		p.Table[1].Rank, p.Table[1].Suit,
		p.Table[2].Rank, p.Table[2].Suit)
}

func PrintPlayerHand(p *Player) {
	fmt.Printf("Player hand: \t")
	for i, v := range p.Hand {
		fmt.Printf("%v: %v %v \t", i+1, v.Rank, v.Suit)
	}
	fmt.Printf("\n")
}
