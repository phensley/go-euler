package euler054

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler054 *.txt

func init() {
	rand.Seed(time.Now().UnixNano())
	euler.Register("054", "Poker hands", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		randomHands(10000)
		pokerOdds(250000)
	}

	player1Wins := 0
	lines := strings.Split(string(rawfiles["p054_poker.txt"]), "\n")
	for _, line := range lines {
		cards := strings.Split(line, " ")
		if len(cards) != 10 {
			continue
		}

		hand1 := parseHand(cards[:5])
		hand2 := parseHand(cards[5:])
		if hand1.Compare(hand2) == 1 {
			player1Wins++
		}
	}

	answer := fmt.Sprintf("%d", player1Wins)
	ctx.SetAnswer(answer)
}

func (h hand) Compare(o *hand) int {
	if h.rank != o.rank {
		if h.rank < o.rank {
			return -1
		}
		return 1
	}

	// If hand ranks are equal, break ties below..
	//
	// The first dimension of the hand index is the count of
	// cards of equal value.  So to compare two sets of quads,
	// we access index 4 and then index 0:
	//  compare(h.index[4][0], h.index[4][0])

	switch h.rank {
	case royalFlush:
		// all 4 possible royal flushes are equal
		return 0

	case fourOfAKind:
		// higher quads wins
		return compare(h.index[4][0], o.index[4][0])

	case threeOfAKind, fullHouse:
		// higher trips wins.
		return compare(h.index[3][0], o.index[3][0])

	case twoPairs:
		// higher pair wins
		r := compare(h.index[2][0], o.index[2][0])
		if r != 0 {
			return r
		}
		// .. or else second pair wins
		r = compare(h.index[2][1], o.index[2][1])
		if r != 0 {
			return r
		}
		// .. or else last high card wins
		return compare(h.index[1][0], o.index[1][0])

	case onePair:
		// Higher pair wins
		r := compare(h.index[2][0], o.index[2][0])
		if r != 0 {
			return r
		}
		// Otherwise remaining high card wins
		for i := 0; i < 3; i++ {
			r = compare(h.index[1][i], o.index[1][i])
			if r != 0 {
				return r
			}
		}

	case straight, straightFlush, flush, noPair:
		// Hand having the higest distinct card wins
		for i := 0; i < 5; i++ {
			r := compare(h.index[1][i], o.index[1][i])
			if r != 0 {
				return r
			}
		}
	}

	// Must be exactly same cards in the hand, differing only by suit.
	// Split the pot..
	return 0
}

func compare(a, b value) int {
	switch {
	case a < b:
		return -1
	case a == b:
		return 0
	default:
		return 1
	}
}
