package util

import (
	"math/rand"
	"time"

	c "goshit/card"
)


// Shuffle creates a random permutation of index values from len(vals) and uses these to re-assign vals positions
func Shuffle(vals []c.Card) []c.Card {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	sDeck := make([]c.Card, len(vals))
	perm := r.Perm(len(vals))
	for i, randIndex := range perm {
		sDeck[i] = vals[randIndex]
	}
	return sDeck
}