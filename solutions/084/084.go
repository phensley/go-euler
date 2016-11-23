package euler084

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("084", "Monopoly odds", solve)
}

var (
	random = rand.New(rand.NewSource(1001))
)

func show(p pos) {
	fmt.Println(name(p), p)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		// At around 4 million iterations the 6-sided solution converges.
		sixIters := 4000000
		fmt.Println("known: ", compute(6, sixIters))
	}

	// The 4-sided solution converges at a much lower number of iterations
	fourIters := 50000
	answer := compute(4, fourIters)
	ctx.SetAnswer(answer)
}

// Monte Carlo simulation of the probability of landing on a given space
// in Monopoly, given N-sided dice and M dice rolls
func compute(diceSides, iterations int) string {
	stats := make(scores, spaceH2+1)
	for p := pos(0); p <= spaceH2; p++ {
		stats[p].p = p
	}

	simulate(diceSides, iterations, stats)

	sort.Sort(sort.Reverse(stats))
	if euler.Verbose {
		for _, s := range stats {
			ratio := (float64(s.score) / float64(iterations)) * 100.0
			fmt.Printf("  %5s  %d   %.5f\n", name(pos(s.p)), s.score, ratio)
		}
	}
	return fmt.Sprintf("%02d%02d%02d", stats[0].p, stats[1].p, stats[2].p)
}

// Simulate N dice rolls and moves in Monopoly
func simulate(diceSides, iterations int, stats []score) int {
	moves := 0
	doubles := 0

	ccCards := shuffleDeck(buildCCDeck())
	chCards := shuffleDeck(buildCHDeck())

	ccIdx := 0
	chIdx := 0
	ccLen := len(ccCards)
	chLen := len(chCards)

	iter := 0

	p := spaceGO
	stats[int(p)].score++

	for iter < iterations {
		iter++

		d1, d2 := roll(diceSides)
		if d1 == d2 {
			doubles++
		} else {
			doubles = 0
		}

		if doubles == 3 {
			p = spaceJAIL
			stats[int(p)].score++
			moves++
			doubles = 0
			continue
		}

		spaces := d1 + d2
		p = move(p, spaces)

		switch p {
		case spaceCH1, spaceCH2, spaceCH3:
			p = chCards[chIdx%chLen](p)
			chIdx++
		}

		switch p {
		case spaceCC1, spaceCC2, spaceCC3:
			p = ccCards[ccIdx%ccLen](p)
			ccIdx++
		}

		if p == spaceG2J {
			p = spaceJAIL
		}

		stats[int(p)].score++
		moves++
	}
	return moves
}

type scores []score

func (s scores) Len() int {
	return len(s)
}

func (s scores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s scores) Less(i, j int) bool {
	return s[i].score <= s[j].score
}

type score struct {
	p     pos
	score int
}

type card func(pos) pos

type ccDeck struct {
	index int
}

func ignoreCard(p pos) pos {
	return p
}

func buildCCDeck() []card {
	return []card{
		func(pos) pos { return spaceGO },
		func(pos) pos { return spaceJAIL },
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
	}
}

func buildCHDeck() []card {
	return []card{
		func(p pos) pos { return spaceGO },
		func(p pos) pos { return spaceJAIL },
		func(p pos) pos { return spaceC1 },
		func(p pos) pos { return spaceE3 },
		func(p pos) pos { return spaceH2 },
		func(p pos) pos { return spaceR1 },
		nextRailway,
		nextRailway,
		nextUtility,
		func(p pos) pos { return p - 3 },
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
		ignoreCard,
	}
}

func shuffleDeck(deck []card) []card {
	r := []card{}
	for _, i := range random.Perm(len(deck)) {
		r = append(r, deck[i])
	}
	return r
}

func nextRailway(p pos) pos {
	switch p {
	case spaceCH1:
		return spaceR2
	case spaceCH2:
		return spaceR3
	}
	return spaceR1
}

func nextUtility(p pos) pos {
	if p == spaceCH2 {
		return spaceU2
	}
	return spaceU1
}

func move(p, spaces pos) pos {
	return (p + spaces) % (spaceH2 + 1)
}

func name(p pos) string {
	return spaceNames[p]
}

func roll(sides int) (pos, pos) {
	d1 := random.Intn(sides) + 1
	d2 := random.Intn(sides) + 1
	return pos(d1), pos(d2)
}

type pos int

var (
	spaceNames = []string{
		"GO", "A1", "CC1", "A2", "T1", "R1", "B1", "CH1", "B2", "B3",
		"JAIL", "C1", "U1", "C2", "C3", "R2", "D1", "CC2", "D2", "D3",
		"FP", "E1", "CH2", "E2", "E3", "R3", "F1", "F2", "U2", "F3",
		"G2J", "G1", "G2", "CC3", "G3", "R4", "CH3", "H1", "T2", "H2",
	}
)

const (
	spaceGO = pos(iota)
	spaceA1
	spaceCC1
	spaceA2
	spaceT1
	spaceR1
	spaceB1
	spaceCH1
	spaceB2
	spaceB3

	spaceJAIL
	spaceC1
	spaceU1
	spaceC2
	spaceC3
	spaceR2
	spaceD1
	spaceCC2
	spaceD2
	spaceD3

	spaceFP
	spaceE1
	spaceCH2
	spaceE2
	spaceE3
	spaceR3
	spaceF1
	spaceF2
	spaceU2
	spaceF3

	spaceG2J
	spaceG1
	spaceG2
	spaceCC3
	spaceG3
	spaceR4
	spaceCH3
	spaceH1
	spaceT2
	spaceH2
)
