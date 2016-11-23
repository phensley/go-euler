package euler093

import (
	"fmt"
	"sort"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("093", "Arithmetic expressions", solve)
}

func solve(ctx *euler.Context) {
	res := calculate()
	answer := fmt.Sprintf("%d%d%d%d", res[0], res[1], res[2], res[3])
	ctx.SetAnswer(answer)
}

type opfunc func(a, b float64) float64

type operator struct {
	sym   string
	apply opfunc
}

type optree func(a, b, c, d float64) float64

var (
	add      = operator{"+", func(a, b float64) float64 { return a + b }}
	subtract = operator{"-", func(a, b float64) float64 { return a - b }}
	multiply = operator{"*", func(a, b float64) float64 { return a * b }}
	divide   = operator{"/", func(a, b float64) float64 { return a / b }}
)

func calculate() []int {
	permutations := euler.NewPermutations([]int{0, 1, 2, 3}).All()
	digits := []int{0, 0, 0, 0}

	trees := buildOperatorTrees()

	longest := 0
	longestDigits := []int{}
	for a := 1; a < 10; a++ {
		for b := a + 1; b < 10; b++ {
			for c := b + 1; c < 10; c++ {
				for d := c + 1; d < 10; d++ {

					res := make(map[int]*struct{})

					for _, perm := range permutations {
						digits[perm[0]] = a
						digits[perm[1]] = b
						digits[perm[2]] = c
						digits[perm[3]] = d

						x, y, w, z := digits[0], digits[1], digits[2], digits[3]
						for _, tree := range trees {
							r := tree(float64(x), float64(y), float64(w), float64(z))

							// Ensure result is an integer
							if r != float64(int(r)) {
								continue
							}
							res[int(r)] = &struct{}{}
						}

					}

					n := sequenceLength(res)
					if n >= longest {
						longest = n
						longestDigits = []int{a, b, c, d}
					}

				}
			}
		}
	}
	return longestDigits
}

func sequenceLength(m map[int]*struct{}) int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	sort.Ints(keys)

	count := 0
	p := -1
	for _, n := range keys {
		if n < 1 {
			continue
		}
		if p != -1 && n != p+1 {
			break
		}
		count++
		p = n
	}
	return count
}

func buildOperatorTrees() []optree {
	indices := euler.NewMixedRadixCounter([]int{4, 4, 4}).All()
	operators := []operator{multiply, divide, add, subtract}

	trees := []optree{}

	for _, idx := range indices {
		ops := []operator{add, add, add}
		for i, j := range idx {
			ops[i] = operators[j]
		}

		oaf, obf, ocf := ops[0].apply, ops[1].apply, ops[2].apply

		// Append 5 operator trees for each permutation
		trees = append(trees, tree1(oaf, obf, ocf))
		trees = append(trees, tree2(oaf, obf, ocf))
		trees = append(trees, tree3(oaf, obf, ocf))
		trees = append(trees, tree4(oaf, obf, ocf))
		trees = append(trees, tree5(oaf, obf, ocf))
	}
	return trees
}

// a(b(cd))
func tree1(oa, ob, oc opfunc) optree {
	return func(a, b, c, d float64) float64 {
		return oa(a, ob(b, oc(c, d)))
	}
}

// (a(bc))d
func tree2(oa, ob, oc opfunc) optree {
	return func(a, b, c, d float64) float64 {
		return oc(oa(a, ob(b, c)), d)
	}
}

// ((ab)c)d
func tree3(oa, ob, oc opfunc) optree {
	return func(a, b, c, d float64) float64 {
		return oc(ob(oa(a, b), c), d)
	}
}

// (ab)(cd)
func tree4(oa, ob, oc opfunc) optree {
	return func(a, b, c, d float64) float64 {
		return ob(oa(a, b), oc(c, d))
	}
}

// a((bc)d)
func tree5(oa, ob, oc opfunc) optree {
	return func(a, b, c, d float64) float64 {
		return oa(a, oc(ob(b, c), d))
	}
}
