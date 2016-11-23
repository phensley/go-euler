package euler

import (
	"math"
)

// Counter ...
type Counter struct {
	e      []int
	length int
	n      int
	m      int
	i      int
}

// NewCounter generates a sequence of integer arrays representing an
// N-digit counter, where each digit takes the value 0-M.
//
// The underlying array is returned from Get() and modified on every
// call to Next() so you must copy it if you plan to cache the sequence.
func NewCounter(n, m int) *Counter {
	e := make([]int, n)
	e[n-1] = -1
	length := int(math.Pow(float64(m), float64(n)))
	return &Counter{e, length, n, m, n - 1}
}

// Next returns true if there are more results available; false if the
// sequence has been completely generated.
func (c *Counter) Next() bool {
	c.e[c.i]++
	if c.e[c.i] < c.m {
		return true
	}

	j := c.i
	for c.e[j] == c.m {
		c.e[j] = 0
		j--
		if j < 0 {
			break
		}
		c.e[j]++
	}
	if j == -1 {
		return false
	}
	return true
}

// Get returns a reference to the underlying array. You must copy this if
// you need to use the returned value after calling Next().
func (c *Counter) Get() []int {
	if c.e[c.i] == -1 {
		panic("Get() called before first call to Next()")
	}
	return c.e
}

// GetCopy returns the result of Get() as a copy
func (c *Counter) GetCopy() []int {
	return CopyIntSlice(c.Get())
}

// Reset ...
func (c *Counter) Reset() {
	c.i = c.n - 1
	for i := 0; i < len(c.e); i++ {
		c.e[i] = 0
	}
	c.e[c.i] = -1
}

// All ...
func (c *Counter) All() [][]int {
	res := make([][]int, c.length)
	i := 0
	for c.Next() {
		res[i] = c.GetCopy()
		i++
	}
	return res
}

// MixedRadixCounter is a counter that generates the full sequence
// of mixed-radix numbers of length M, where each digit can have a
// different base N.
//  The sequence length is the product of the M bases.
type MixedRadixCounter struct {
	m      []int
	length int
	n      int
	i      int
	e      []int
}

// NewMixedRadixCounter ...
// TODO: clean up
func NewMixedRadixCounter(m []int) *MixedRadixCounter {
	e := make([]int, len(m))
	i := len(m) - 1
	e[i] = -1
	length := 1
	for _, n := range m {
		length *= n
	}
	return &MixedRadixCounter{m, length, len(m), i, e}
}

// Next ..
func (m *MixedRadixCounter) Next() bool {
	k := m.i
	for m.e[k] == m.m[k]-1 {
		m.e[k] = 0
		k--
		if k < 0 {
			return false
		}
	}
	m.e[k]++
	return true
}

// Get ...
func (m *MixedRadixCounter) Get() []int {
	return m.e
}

// GetCopy ..
func (m *MixedRadixCounter) GetCopy() []int {
	return CopyIntSlice(m.e)
}

// Reset ...
func (m *MixedRadixCounter) Reset() {
	for i := 0; i < m.n; i++ {
		m.e[i] = 0
	}
	m.i = m.n - 1
	m.e[m.i] = -1
}

// All ...
func (m *MixedRadixCounter) All() [][]int {
	res := make([][]int, m.length)
	i := 0
	for m.Next() {
		res[i] = m.GetCopy()
		i++
	}
	return res
}

// Rotations stores state for rotations
type Rotations struct {
	n   []int
	i   int
	lim int
}

// NewRotations generates the rotations of n
func NewRotations(n []int) *Rotations {
	return &Rotations{n, -1, len(n) - 1}
}

// Next returns true if another rotation is available
func (r *Rotations) Next() bool {
	r.i++
	return r.i < len(r.n)
}

// Get gets the current rotation state
func (r *Rotations) Get() []int {
	if r.i > 0 && r.i <= r.lim {
		t := r.n[r.lim]
		for i := r.lim; i > 0; i-- {
			r.n[i] = r.n[i-1]
		}
		r.n[0] = t
	}
	return r.n
}

// GetCopy returns a copy of the result of Get()
func (r *Rotations) GetCopy() []int {
	return CopyIntSlice(r.Get())
}

// All ..
func (r *Rotations) All() [][]int {
	if r.i > -1 {
		panic("sequence already read from")
	}
	res := make([][]int, len(r.n))
	i := 0
	for r.Next() {
		res[i] = r.GetCopy()
		i++
	}
	return res
}

// Permutations generates all permutations of the integer
// array in lexicographic order.
type Permutations struct {
	n     []int
	e     []int
	k     int
	start bool
}

// NewPermutations ...
func NewPermutations(n []int) *Permutations {
	assertAscending(n)
	e := make([]int, len(n))
	copy(e, n)
	return &Permutations{n, e, 0, false}
}

// Next ..
func (o *Permutations) Next() bool {
	if !o.start {
		o.start = true
		return true
	}

	k := findK(o.e)
	if k == -1 {
		return false
	}
	i := findI(k, o.e)
	o.e[k], o.e[i] = o.e[i], o.e[k]
	ReverseIntSlice(o.e[k+1:])
	return true
}

// Get ..
func (o *Permutations) Get() []int {
	return o.e
}

// GetCopy ...
func (o *Permutations) GetCopy() []int {
	return CopyIntSlice(o.e)
}

// All ...
func (o *Permutations) All() [][]int {
	if o.start {
		o.Reset()
	}
	res := [][]int{}
	for o.Next() {
		res = append(res, o.GetCopy())
	}
	return res
}

// Reset ...
func (o *Permutations) Reset() {
	copy(o.e, o.n)
	o.k = 0
	o.start = false
}

func findK(e []int) int {
	for k := len(e) - 2; k >= 0; k-- {
		if e[k] < e[k+1] {
			return k
		}
	}
	return -1
}

func findI(k int, e []int) int {
	for i := len(e) - 1; i > k; i-- {
		if e[k] < e[i] {
			return i
		}
	}
	return -1
}

// PartialPermutations generates the K permutations of N
type PartialPermutations struct {
	n      []int
	length int
	k      int
	e      []int
	start  bool
}

// NewPartialPermutations ...
func NewPartialPermutations(n []int, k int) *PartialPermutations {
	length := len(n)
	e := CopyIntSlice(n)
	return &PartialPermutations{n, length, k, e, false}
}

// Next generates the next partial permutation
func (p *PartialPermutations) Next() bool {
	if !p.start {
		p.start = true
		return true
	}
	n := p.length
	i := 0
	j := 0
	e := p.k - 1
	if p.k < n {
		j = p.k
		for j < n && p.e[e] >= p.e[j] {
			j++
		}
	}

	if p.k < n && j < n {
		p.e[e], p.e[j] = p.e[j], p.e[e]
	} else {
		if p.k < n {
			reverse(p.e[p.k:], n-p.k)
		}

		i = e - 1
		for i >= 0 && p.e[i] >= p.e[i+1] {
			i--
		}

		if i < 0 {
			return false
		}

		j = n - 1
		for j > i && p.e[i] >= p.e[j] {
			j--
		}
		p.e[i], p.e[j] = p.e[j], p.e[i]
		reverse(p.e[i+1:], n-i-1)
	}
	return true
}

// Get ..
func (p *PartialPermutations) Get() []int {
	return p.e[:p.k]
}

// GetCopy ..
func (p *PartialPermutations) GetCopy() []int {
	return CopyIntSlice(p.e[:p.k])
}

// All ...
func (p *PartialPermutations) All() [][]int {
	if p.start {
		p.Reset()
	}
	res := [][]int{}
	for p.Next() {
		res = append(res, p.GetCopy())
	}
	return res
}

// Reset ...
func (p *PartialPermutations) Reset() {
	copy(p.e, p.n)
	p.k = 0
	p.start = false
}

// Combinations  ...
func Combinations(iterable []int, r int) [][]int {
	pool := iterable
	n := len(pool)

	combs := [][]int{}
	if r > n {
		return combs
	}
	indices := []int{}
	for i := 0; i < r; i++ {
		indices = append(indices, i)
	}

	tuple := []int{}
	for _, i := range indices {
		tuple = append(tuple, pool[i])
	}
	combs = append(combs, tuple)

	for {
		i := r - 1
		for ; i >= 0; i-- {
			if indices[i] != i+n-r {
				break
			}
		}
		if i < 0 {
			return combs
		}
		indices[i]++
		for j := i + 1; j < r; j++ {
			indices[j] = indices[j-1] + 1
		}
		tuple2 := []int{}
		for _, i := range indices {
			tuple2 = append(tuple2, pool[i])
		}
		combs = append(combs, tuple2)
	}
}

func reverse(a []int, length int) {
	i := 0
	j := length - 1
	for i < j {
		a[i], a[j] = a[j], a[i]
		i++
		j--
	}
}

func assertAscending(e []int) {
	if len(e) <= 1 {
		return
	}
	for i := 1; i < len(e); i++ {
		if e[i] < e[i-1] {
			panic("elements of sequence are not in ascending order!")
		}
	}
}
