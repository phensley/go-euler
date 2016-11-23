package euler042

import (
	"fmt"
	"strings"

	"github.com/phensley/go-euler"
)

//go:generate ../embed euler042 *.txt

func init() {
	euler.Register("042", "Coded triangle numbers", solve)
}

func solve(ctx *euler.Context) {
	lines := strings.Split(string(rawfiles["p042_words.txt"]), ",")
	t := &triangles{0, make(map[int]*struct{}), 1}

	count := 0
	for i := 0; i < len(lines); i++ {
		word := lines[i]
		word = word[1 : len(word)-1]
		sum := 0
		for _, c := range word {
			v := int(byte(c)-byte('A')) + 1
			sum += v
		}
		if t.check(sum) {
			count++
		}
	}

	answer := fmt.Sprintf("%d", count)
	ctx.SetAnswer(answer)
}

type triangles struct {
	max int
	set map[int]*struct{}
	n   int
}

func (t *triangles) check(n int) bool {
	// Generate more triangle numbers on demand
	for t.max < n {
		r := int(0.5 * float64(t.n) * (float64(t.n) + 1))
		t.set[r] = &struct{}{}
		t.n++
		t.max = r
	}
	return t.set[n] != nil
}
