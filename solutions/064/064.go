package euler064

import (
	"fmt"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("064", "Odd period square roots", solve)
}

func solve(ctx *euler.Context) {
	limit := 10000
	s := solution{integerRoots(limit)}
	odds := 0
	for n := 2; n <= limit; n++ {
		r := s.compute(n)
		if len(r.d)&1 != 0 {
			odds++
		}
	}
	answer := fmt.Sprintf("%d", odds)
	ctx.SetAnswer(answer)
}

type solution struct {
	roots map[int]int
}

type cf struct {
	i int
	d []int
}

func (s *solution) compute(n int) *cf {
	r := s.roots[n]
	d := []int{}
	a := 1
	b := r
	for {
		a = (n - b*b) / a
		if a == 0 {
			break
		}
		d = append(d, (r+b)/a)
		b = r - (r+b)%a
		if a <= 1 {
			break
		}
	}
	return &cf{r, d}
}

func integerRoots(limit int) map[int]int {
	res := map[int]int{}
	a := 2
	i := 0
	for i <= limit {
		r := a * a
		for i < r {
			res[i] = a - 1
			i++
		}
		a++
	}
	return res
}
