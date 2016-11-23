package euler052

import (
	"fmt"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("052", "Permuted multiples", solve)
}

func solve(ctx *euler.Context) {
	g := generator{1, 0, nil}

outer:
	for g.next() {
		n, digits := g.get()
		for f := 3; f <= 6; f++ {
			d := euler.IntToDigits(n * f)
			sort.Ints(d)
			if !equal(digits, d) {
				continue outer
			}
		}

		if euler.Verbose {
			for f := 1; f <= 6; f++ {
				fmt.Println(f * n)
			}
		}

		answer := fmt.Sprintf("%d", n)
		ctx.SetAnswer(answer)
		break
	}
}

type generator struct {
	i int
	x int
	y []int
}

// Generate next candidate pair N and 2N having the same digits
func (g *generator) next() bool {
	for g.i <= 1000000 {
		j := 2 * g.i
		di := euler.IntToDigits(g.i)
		dj := euler.IntToDigits(j)
		sort.Ints(di)
		if di[0] != 0 {
			sort.Ints(dj)
			if dj[0] != 0 && equal(di, dj) {
				g.x = g.i
				g.y = di
				g.i++
				return true
			}
		}
		g.i++
	}
	return false
}

func (g *generator) get() (int, []int) {
	return g.x, g.y
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
