package main

import (
	"fmt"

	"goshit/util"
	"goshit/card"
)



func init() {}

func main() {
	rawDeck := card.CreateRawDeck()
	deck := util.Shuffle(rawDeck)
	for i, c := range deck {
		fmt.Println(i, c)
	}
}
