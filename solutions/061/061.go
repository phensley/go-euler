package euler061

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("061", "Cyclical figurate numbers", solve)
}

func solve(ctx *euler.Context) {
	// Min/max for the 4-digit numbers we're interested in.
	min := uint64(1000)
	max := uint64(10000)

	// A linkset for each sequence of numbers we're working with
	triangle := makeLinkset(euler.Triangle(min, max))
	square := makeLinkset(euler.Square(min, max))
	pentagonal := makeLinkset(euler.Pentagonal(min, max))
	hexagonal := makeLinkset(euler.Hexagonal(min, max))
	heptagonal := makeLinkset(euler.Heptagonal(min, max))
	octagonal := makeLinkset(euler.Octagonal(min, max))

	// First, confirm we solved the known
	sets := []*linkset{square, pentagonal, triangle}
	indices := []int{0, 1, 2}
	known := calculate(indices, sets)
	if known != "19291" {
		fmt.Println("FAIL: known doesn't match: ", known)
		return
	}

	// Now add in more sets and solve the unknown
	sets = []*linkset{square, pentagonal, triangle, hexagonal, heptagonal, octagonal}
	indices = []int{0, 1, 2, 3, 4, 5}
	answer := calculate(indices, sets)
	ctx.SetAnswer(answer)
}

// A link in the chain: a 4-digit number, with its lead and tail digits.
type link struct {
	n    uint64
	lead uint64
	tail uint64
}

// Links for a given set (triangle, square, etc) with lookup tables
// for particular lead and tail digits.
type linkset struct {
	links   []link
	leadMap map[uint64][]link
	tailMap map[uint64][]link
}

func makeLinks(b *euler.Bitstring) []link {
	r := make([]link, b.Count())
	iter := b.OneIterator()
	i := 0
	for iter.Next() {
		n := iter.Get()
		r[i] = link{n, n / 100, n % 100}
		i++
	}
	return r
}

func makeLinkset(b *euler.Bitstring) *linkset {
	links := makeLinks(b)
	leadMap := map[uint64][]link{}
	tailMap := map[uint64][]link{}
	for _, link := range links {
		leadMap[link.lead] = append(leadMap[link.lead], link)
		tailMap[link.tail] = append(tailMap[link.tail], link)
	}
	return &linkset{links, leadMap, tailMap}
}

// Recursively search through the linksets until we hit bottom.
// Once we hit the last set, check the matches against the first link.
// If they match we found the answer.
func searchNext(first link, links []link, rest []*linkset) []link {
	set := rest[0]
	rest = rest[1:]
	last := len(rest) == 0
	for _, lk := range links {
		matches := set.leadMap[lk.tail]
		if matches == nil {
			continue
		}
		if last {
			for _, m := range matches {
				if m.tail == first.lead {
					return []link{lk, m}
				}
			}
		} else {
			r := searchNext(first, matches, rest)
			if r != nil {
				r = append([]link{lk}, r...)
				return r
			}
		}
	}
	return nil
}

// Start a search using the first linkset.
func search(sets []*linkset) []link {
	set := sets[0]
	next := sets[1]
	rest := sets[2:]
	for _, first := range set.links {
		matches := next.leadMap[first.tail]
		if matches == nil {
			continue
		}
		r := searchNext(first, matches, rest)
		if r != nil {
			r = append([]link{first}, r...)
			return r
		}
	}
	return nil
}

// Reorder the sets according to the indices
func permuteSets(indices []int, sets []*linkset) []*linkset {
	r := make([]*linkset, len(sets))
	for i, j := range indices {
		r[i] = sets[j]
	}
	return r
}

func calculate(indices []int, sets []*linkset) string {
	perms := euler.NewPermutations(indices)
	for perms.Next() {
		perm := perms.Get()
		set := permuteSets(perm, sets)
		answer := search(set)
		if answer != nil {
			sum := uint64(0)
			for _, lk := range answer {
				sum += lk.n
			}
			return fmt.Sprintf("%d", sum)
		}
	}
	return ""
}
