package euler086

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("086", "Cuboid route", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		m, c := calculateDirect(1975)
		fmt.Printf("known m=%d produced %d cuboids\n", m, c)

		m, c = calculateDirect(2000)
		fmt.Printf("known m=%d produced %d cuboids\n", m, c)
	}

	goal := 1000000
	m, c := calculateDirect(goal)
	if euler.Verbose {
		fmt.Printf("known m=%d produced %d cuboids\n", m, c)

		// Sanity check using the Pythagorean triple method
		maxSide := uint64(4000)
		triples := generateTriples(maxSide)
		m2, c2 := calculateTriples(goal, triples)
		if m != m2 || c != c2 {
			fmt.Println("error in confirmation using Pythagorean triples: ", m2, c2)
		}
	}

	answer := fmt.Sprintf("%d", m)
	ctx.SetAnswer(answer)
}

// Our cuboid is formed by 2 values: a is the base length, and b is the sum of
// the lengths of the remaining 2 sides.
// Instead of pre-generating Pythagorean triples we discover them by finding
// values of (a, b) which produce an integral hypotenuse (c). We gradually
// increase our triangle's base side by 1, counting all cuboids aligned with
// our integer right triangle.
func calculateDirect(goal int) (int, int) {
	count := 0
	a := 2
	for {
		a2 := a * a
		h := a
		for b := 3; b < 2*a; b++ {
			c := b*b + a2

			// Increase until hypotenuse is integral or overshoots
			for h*h < c {
				h++
			}
			if h*h == c {
				if b <= a {
					v := b / 2
					count += v
				} else {
					v := 1 + (a - (b+1)/2)
					count += v
				}
				h++
			}
		}
		if count >= goal {
			break
		}
		a++
	}
	return a, count
}

// Calculate using array of Pythagorean triples. Initially solved using this
// technique. Downside is we compute the triples in advance but don't know
// the maximum side length to use, so requires some trial and error.
func calculateTriples(goal int, triples map[int][]*euler.Triple) (int, int) {
	count := 0
	m := 3
	for {
		trips := triples[m]
		m2 := 2 * m
		if trips != nil {
			for _, t := range trips {
				if t.B <= t.A {
					v := int(t.B) / 2
					count += v
				} else if int(t.B) <= m2 {
					v := 1 + (m - (int(t.B)+1)/2)
					count += v
				}
			}
		}
		if count >= goal {
			break
		}
		m++
	}
	return m, count
}

// Generate triples meeting our criteria
func generateTriples(maxSide uint64) map[int][]*euler.Triple {
	m := map[int][]*euler.Triple{}
	limiter := func(t *euler.Triple) bool {
		return t.A > maxSide || t.B > maxSide
	}
	for trip := range euler.PythagoreanTriplesBerggren(limiter) {
		f := uint64(1)
		tt := trip.Multiply(1)
		for tt.A <= maxSide || tt.B <= maxSide {
			a := int(tt.A)
			b := int(tt.B)

			m[a] = append(m[a], tt)
			m[b] = append(m[b], euler.NewTriple(tt.B, tt.A, tt.C))
			f++
			tt = trip.Multiply(f)
		}
	}
	return m
}
