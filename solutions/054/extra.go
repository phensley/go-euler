package euler054

import (
	"fmt"
	"sort"
)

// Some extra credit:
// - confirm poker odds by generating random hands
// - play random hands against one another

func pokerOdds(limit int) {
	ranks := make(map[rank]int)
	suits := make(map[suit]int)
	values := make(map[value]int)
	i := 0

	fmt.Printf("Confirming poker hand odds, generating %d random hands ...\n", limit)

	d := newDeck()

	for i < limit {
		d.shuffle()
		h := newHand(d.draw(), d.draw(), d.draw(), d.draw(), d.draw())
		ranks[h.rank]++

		for _, c := range h.cards {
			suits[c.suit]++
			values[c.value]++
		}
		i++
	}

	fmt.Println("Suit counts:")
	for s, c := range suits {
		fmt.Printf("suit %s  %10d\n", s, c)
	}
	fmt.Println()

	valueKeys := valueslice{}
	for v := range values {
		valueKeys = append(valueKeys, v)
	}
	sort.Sort(valueKeys)

	fmt.Println("Value counts:")
	for _, v := range valueKeys {
		fmt.Printf("%s  %10d\n", v, values[v])
	}
	fmt.Println()

	rankKeys := rankslice{}
	for k := range ranks {
		rankKeys = append(rankKeys, k)
	}
	sort.Sort(rankKeys)

	fmt.Println("Hand frequencies:")
	for _, k := range rankKeys {
		c := ranks[k]
		r := (float64(c) / float64(limit)) * 100
		fmt.Printf("%8.5f%%   %10d   %s\n", r, c, rank(k))
	}
	fmt.Println()
}

func randomHands(limit int) {
	i := 0
	d := newDeck()

	for i < limit {
		d.shuffle()
		h1 := d.deal()
		h2 := d.deal()

		switch h1.Compare(h2) {
		case -1:
			fmt.Println(h2, "   beats   ", h1)
		case 0:
			fmt.Println(h1, "   draws   ", h2)
		case 1:
			fmt.Println(h1, "   beats   ", h2)
		}

		i++
	}
}
