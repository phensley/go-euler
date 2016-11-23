package euler068

import (
	"fmt"
	"math"
	"math/big"
	"sort"
	"strconv"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("068", "Magic 5-gon ring", solve)
}

func solve(ctx *euler.Context) {
	if euler.Verbose {
		// Largest 9-digit number for a 3-gon
		calculate(3, 9, []int{1, 2, 3, 4, 5, 6})
	}

	// Largest 16-digit number for a 5-gon
	max := calculate(5, 16, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	answer := fmt.Sprintf("%s", max)
	ctx.SetAnswer(answer)
}

func calculate(size, length int, digits []int) *big.Int {
	if euler.Verbose {
		fmt.Printf("calculating %d-gon using digits: %v\n", size, digits)
	}

	// Generate permutations of digits of size 3
	sets := euler.NewPartialPermutations(digits, 3).All()

	// Map nodes by their midpoint. We'll use these to resolve links
	links := map[int][][]int{}
	for _, p := range sets {
		m := p[1]
		links[m] = append(links[m], p)
	}

	// Track distinct rings only
	seen := make(map[string]*struct{})

	// Store big.Int results
	results := bigintlist{}

	// Generate all rings of target size, using the integer sets and link map.
	g := newGenerator(size, sets, links)
	for r := range g.generate() {
		r.rotate()
		s := r.String()
		if seen[s] != nil {
			continue
		}

		seen[s] = &struct{}{}

		// Build a string representation of the digits
		num := r.digits()
		sd := fmt.Sprintf("%s", num)
		if len(sd) == length {
			results = append(results, num)
		}

		if euler.Verbose {
			// Display human-readable form of N-gon
			fmt.Printf("  %3d  %s\n", r.sum, s)
		}
	}

	sort.Sort(sort.Reverse(results))
	if euler.Verbose {
		fmt.Printf("maximum string of length %d: %s\n\n", length, results[0])
	}
	return results[0]
}

type ring struct {
	sum   int
	size  int
	lines [][]int
}

func (r *ring) rotate() {
	// Find lowest starting digit of the ring
	min := math.MaxInt32
	idx := -1
	for i := 0; i < r.size; i++ {
		if r.lines[i][0] < min {
			min = r.lines[i][0]
			idx = i
		}
	}

	// Rotate ring so the line with lowest digit is first
	rotated := make([][]int, len(r.lines))
	for i := 0; i < r.size; i++ {
		j := (i + idx) % r.size
		rotated[i] = r.lines[j]
	}
	r.lines = rotated
}

func (r *ring) String() string {
	s := ""
	for i := 0; i < len(r.lines); i++ {
		if i > 0 {
			s += "; "
		}
		a := r.lines[i]
		s += fmt.Sprintf("%d,%d,%d", a[0], a[1], a[2])
	}
	return s
}

func (r *ring) digits() *big.Int {
	s := ""
	for i := 0; i < len(r.lines); i++ {
		a := r.lines[i]
		for j := 0; j < 3; j++ {
			s += strconv.Itoa(a[j])
		}
	}
	n, _ := big.NewInt(0).SetString(s, 10)
	return n
}

type generator struct {
	size  int
	sets  [][]int
	links map[int][][]int
}

func newGenerator(size int, sets [][]int, links map[int][][]int) *generator {
	return &generator{size, sets, links}
}

func (g *generator) generate() chan *ring {

	// Start a background recursive search. For each starting set,
	// try all links until a ring of desired size is reached.
	ch := make(chan *ring, 0)
	go func() {
		// For each possible starting set, search for complete rings
		for _, s := range g.sets {
			g.search(ch, sumInt(s...), [][]int{s})
		}
		close(ch)
	}()
	return ch
}

// Recursively search for complete rings starting from this set.
func (g *generator) search(ch chan *ring, sum int, set [][]int) {
	last := len(set) - 1
	p := set[last]
	e := p[2]
	for _, n := range g.links[e] {
		if sumInt(n...) != sum {
			continue
		}

		// Asset that start/end digits differ between all pairs.
		fail := false
		d, f := n[0], n[2]
		for i := 0; i < len(set); i++ {
			r := set[i]
			a, c := r[0], r[2]
			if a == d || a == f || c == d || c == f {
				fail = true
				break
			}
		}

		// See if the ring joins up at the end
		if !fail && len(set)+1 == g.size {
			if n[2] != set[0][1] {
				fail = true
			}
		}

		if !fail {
			// We found at least part of a ring.
			nring := make([][]int, len(set))
			copy(nring, set)
			nring = append(nring, n)
			if len(nring) < g.size {
				// Ring incomplete, search further.
				g.search(ch, sum, nring)
			} else {
				// Ring found, emit it
				ch <- &ring{sum, g.size, nring}
			}
		}
	}
}

type bigintlist []*big.Int

func (b bigintlist) Len() int {
	return len(b)
}

func (b bigintlist) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b bigintlist) Less(i, j int) bool {
	return b[i].Cmp(b[j]) < 0
}

func sumInt(a ...int) int {
	r := 0
	for i := 0; i < len(a); i++ {
		r += a[i]
	}
	return r
}
