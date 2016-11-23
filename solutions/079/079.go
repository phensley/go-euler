package euler079

import (
	"fmt"
	"sort"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler079 *.txt

func init() {
	euler.Register("079", "Passcode derivation", solve)
}

func solve(ctx *euler.Context) {
	data := string(rawfiles["p079_keylog.txt"])
	lines := strings.Split(data, "\n")

	s := solution{make(map[int]map[int]*struct{})}

	// Parse the 3-digit sequences out of each line
	for _, d := range lines {
		if len(d) == 0 {
			continue
		}
		c0, c1, c2 := int(d[0]-'0'), int(d[1]-'0'), int(d[2]-'0')

		// Map each digit to the one that follows it
		s.set(c0, c1)
		s.set(c0, c2)
		s.set(c1, c2)
	}
	answer := s.solve()
	ctx.SetAnswer(answer)
}

type solution struct {
	m map[int]map[int]*struct{}
}

func (s *solution) set(a, b int) {
	if s.m[a] == nil {
		s.m[a] = make(map[int]*struct{})
	}
	if s.m[b] == nil {
		s.m[b] = make(map[int]*struct{})
	}
	s.m[a][b] = &struct{}{}
}

// For each digit, examine the length of the chain of following digits
// Sort by the length
func (s *solution) solve() string {
	// Build a list of pairs of {digit, length-of-following-chain}
	digits := pairlist{}
	for d, set := range s.m {
		digits = append(digits, pair{d, len(set)})
	}

	// Reverse sort by chain length
	sort.Sort(sort.Reverse(digits))

	// Build the shortest passcode
	code := ""
	for _, d := range digits {
		code += fmt.Sprintf("%d", d.n)
	}
	return code
}

type pair struct {
	n      int
	length int
}

type pairlist []pair

func (a pairlist) Len() int           { return len(a) }
func (a pairlist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a pairlist) Less(i, j int) bool { return a[i].length < a[j].length }
