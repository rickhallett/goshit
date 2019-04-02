package util

import (
	"encoding/json"
	"fmt"
	"goshit/player"
)

func PrettyPrint(message ...interface{}) {
	for _, m := range message {
		jsonMessage, _ := json.MarshalIndent(m, "", "  ")
		fmt.Println(string(jsonMessage))
	}
}

func PrintPlayerCards(p *player.Player) {
	fmt.Printf("Player Hand: \t1: %v %v \t2: %v %v \t3: %v %v \n" +
		"Player Table: \t1: %v %v \t2: %v %v \t3: %v %v \n" +
		"Player Blind: \t1: ? ? \t\t2: ? ? \t\t3: ? ?\n",
		p.Hand[0].Rank, p.Hand[0].Suit,
		p.Hand[1].Rank, p.Hand[1].Suit,
		p.Hand[2].Rank, p.Hand[2].Suit,
		p.Table[0].Rank, p.Table[0].Suit,
		p.Table[1].Rank, p.Table[1].Suit,
		p.Table[2].Rank, p.Table[2].Suit)
}
