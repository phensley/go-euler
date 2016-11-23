package euler088

import (
	"fmt"
	"math"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("088", "Product-sum numbers", solve)
}

func solve(ctx *euler.Context) {
	unique := make(map[int]*struct{})
	limit := 24000

	// Maximum power of 2 to reach our limit
	max := int(math.Log2(float64(limit)))

	// Generate all possible product sequences up to max of N digits
	seq := sequencelist{}
	for n := 2; n <= max; n++ {
		s := generate(limit, n)
		seq = append(seq, s...)
	}

	sort.Sort(seq)

	// Find the product sums
	start := 0
	for k := 2; k <= 12000; k++ {
		found := false
		for i := start; i < len(seq); i++ {
			s := seq[i]
			if k > s.prod {
				// Skip over impossible solutions
				start = i
				continue
			}

			// Check if product == sum
			x := (k - s.length) + s.sum
			if x == s.prod {
				unique[x] = &struct{}{}
				if euler.Verbose {
					fmt.Printf("found: k=%d  %v\n", k, s.n)
				}
				found = true
				break
			}
		}
		if !found {
			panic(fmt.Sprintf("failed to find k=%d\n", k))
		}
	}

	total := 0
	for k := range unique {
		total += k
	}

	answer := fmt.Sprintf("%d", total)
	ctx.SetAnswer(answer)
}

type sequence struct {
	prod   int
	sum    int
	length int
	n      []int
}

func (s sequence) String() string {
	return fmt.Sprintf("{prod=%d sum=%d len=%d n=%v}", s.prod, s.sum, s.length, s.n)
}

type sequencelist []*sequence

func (s sequencelist) Len() int { return len(s) }

func (s sequencelist) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func (s sequencelist) Less(i, j int) bool {
	return s[i].prod < s[j].prod
}

// Generate all sequences of natural numbers of given length whose
// product is <= given limit.
func generate(limit, length int) []*sequence {
	res := []*sequence{}
	e := make([]int, length)
	for i := 0; i < length; i++ {
		e[i] = 2
	}

	i := 0
	for {
		prod := product(e)
		if prod <= limit {
			s := &sequence{prod, sum(e), length, euler.CopyIntSlice(e)}
			res = append(res, s)
			e[i]++
		} else {
			i++
			for i < length && e[i] >= e[i-1] {
				i++
			}
			if i >= length {
				// Generated all possible sequences
				break
			}

			e[i]++
			for k := 0; k < i; k++ {
				e[k] = e[i]
			}
			i = 0
		}
	}
	return res
}

func sum(ns []int) int {
	s := 0
	for _, n := range ns {
		s += n
	}
	return s
}

func product(ns []int) int {
	p := 1
	for _, n := range ns {
		p *= n
	}
	return p
}
