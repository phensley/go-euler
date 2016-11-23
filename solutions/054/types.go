package euler054

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
)

type value int

type valueslice []value

const (
	two value = iota
	three
	four
	five
	six
	seven
	eight
	nine
	ten
	jack
	queen
	king
	ace
)

var (
	valueStrings = map[value]string{
		two:   "2",
		three: "3",
		four:  "4",
		five:  "5",
		six:   "6",
		seven: "7",
		eight: "8",
		nine:  "9",
		ten:   "T",
		jack:  "J",
		queen: "Q",
		king:  "K",
		ace:   "A",
	}
)

func (v value) String() string {
	return valueStrings[v]
}

func parseValue(v byte) value {
	switch v {
	case 'T':
		return ten
	case 'J':
		return jack
	case 'Q':
		return queen
	case 'K':
		return king
	case 'A':
		return ace
	}
	r := int(v - '2')
	if r < 0 || r > int(nine) {
		panic("invalid card value " + string(v))
	}
	return value(r)
}

func (v valueslice) Len() int { return len(v) }

func (v valueslice) Swap(i, j int) { v[i], v[j] = v[j], v[i] }

func (v valueslice) Less(i, j int) bool { return v[i] < v[j] }

type suit byte

const (
	spade   suit = 'S'
	diamond suit = 'D'
	club    suit = 'C'
	heart   suit = 'H'
)

func (s suit) String() string {
	return string(s)
}

func parseSuit(v byte) suit {
	switch v {
	case 'D':
		return diamond
	case 'S':
		return spade
	case 'H':
		return heart
	case 'C':
		return club
	default:
		panic("invalid card suit " + string(v))
	}
}

type rank int

type rankslice []rank

const (
	noPair rank = iota
	onePair
	twoPairs
	threeOfAKind
	straight
	flush
	fullHouse
	fourOfAKind
	straightFlush
	royalFlush
)

func (r rank) String() string {
	switch r {
	case noPair:
		return "no pair"
	case onePair:
		return "one pair"
	case twoPairs:
		return "two pairs"
	case threeOfAKind:
		return "three of a kind"
	case straight:
		return "straight"
	case flush:
		return "flush"
	case fullHouse:
		return "full house"
	case fourOfAKind:
		return "fourOfAKind"
	case straightFlush:
		return "straightFlush"
	case royalFlush:
		return "royalFlush"
	}

	fmt.Printf("invalid rank!  %#v", r)
	os.Exit(1)
	return ""
}

func (r rankslice) Len() int { return len(r) }

func (r rankslice) Swap(i, j int) { r[i], r[j] = r[j], r[i] }

func (r rankslice) Less(i, j int) bool { return r[i] < r[j] }

type card struct {
	value value
	suit  suit
}

func (c card) String() string {
	return fmt.Sprintf("%s%s", c.value, c.suit)
}

func parseCard(raw string) *card {
	if len(raw) != 2 {
		log.Fatalln("card representations must be 2 characters", raw)
	}

	return &card{
		value: parseValue(raw[0]),
		suit:  parseSuit(raw[1]),
	}
}

type hand struct {
	cards []*card
	rank  rank

	// Count cards by value
	counts map[value]int

	// Orders cards by rank, to more easily compare hands
	index map[int]valueslice
}

func (h hand) Len() int { return len(h.cards) }

func (h hand) Swap(i, j int) { h.cards[i], h.cards[j] = h.cards[j], h.cards[i] }

func (h hand) Less(i, j int) bool { return h.cards[i].value < h.cards[j].value }

func (h hand) String() string {
	c := h.cards
	return fmt.Sprintf("%s %s %s %s %s (%s)", c[0], c[1], c[2], c[3], c[4], h.rank)
}

func parseHand(raw []string) *hand {
	if len(raw) != 5 {
		log.Fatalln("hands must have 5 cards", raw)
	}

	return newHand(
		parseCard(raw[0]),
		parseCard(raw[1]),
		parseCard(raw[2]),
		parseCard(raw[3]),
		parseCard(raw[4]),
	)
}

func newHand(cards ...*card) *hand {
	if len(cards) != 5 {
		log.Fatalln("hands must have 5 cards", cards)
	}
	h := &hand{
		cards,
		noPair,
		make(map[value]int),
		map[int]valueslice{},
	}

	// Accumulate counts by card value, then produce an index for faster comparison
	for _, c := range h.cards {
		h.counts[c.value]++
	}
	for v, c := range h.counts {
		h.index[c] = append(h.index[c], v)
	}
	for i := 1; i <= 4; i++ {
		if h.index[i] != nil {
			sort.Sort(sort.Reverse(h.index[i]))
		}
	}

	h.rank = rankOf(h)
	return h
}

func rankOf(h *hand) rank {
	sort.Sort(h)
	c := h.cards

	isFlush := c[0].suit == c[1].suit &&
		c[0].suit == c[2].suit &&
		c[0].suit == c[3].suit &&
		c[0].suit == c[4].suit

	c0 := int(c[0].value)
	c1 := int(c[1].value)
	c2 := int(c[2].value)
	c3 := int(c[3].value)
	c4 := int(c[4].value)

	isStraight := c1 == (c0+1) && c2 == (c0+2) && c3 == (c0+3) && c4 == (c0+4)

	switch {
	case isFlush && isStraight && c[4].value == ace:
		return royalFlush
	case isFlush && isStraight:
		return straightFlush
	case isFlush:
		return flush
	case isStraight:
		return straight
	}

	// Default is to have no cards matching.
	r := noPair
	for _, count := range h.counts {
		switch count {
		case 4:
			r = fourOfAKind

		case 3:
			if r == onePair {
				r = fullHouse

			} else {
				r = threeOfAKind
			}

		case 2:
			switch r {
			case threeOfAKind:
				r = fullHouse

			case onePair:
				r = twoPairs

			default:
				r = onePair
			}
		}
	}
	return r
}

type deck struct {
	cards    []card
	shuffled []card
}

func newDeck() *deck {
	cards := make([]card, 52)
	i := 0
	for v := two; v <= ace; v++ {
		for _, s := range []suit{club, diamond, heart, spade} {
			cards[i] = card{v, s}
			i++
		}
	}
	return &deck{cards, nil}
}

func (d *deck) shuffle() {
	d.shuffled = make([]card, 52)
	copy(d.shuffled, d.cards)
	shuffle(&d.shuffled)
}

func (d *deck) draw() *card {
	if d.shuffled == nil || len(d.shuffled) == 0 {
		return nil
	}
	c := d.shuffled[0]
	d.shuffled = d.shuffled[1:]
	return &c
}

func (d *deck) deal() *hand {
	return newHand(d.draw(), d.draw(), d.draw(), d.draw(), d.draw())
}

func shuffle(cards *[]card) {
	s := *cards
	for i := range s {
		j := rand.Intn(i + 1)
		s[i], s[j] = s[j], s[i]
	}
	*cards = s
}
