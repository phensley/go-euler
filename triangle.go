package euler

import "fmt"

// Triple ...
type Triple struct {
	A   uint64
	B   uint64
	C   uint64
	Sum uint64
}

// Multiply ..
func (t *Triple) Multiply(n uint64) *Triple {
	return NewTriple(t.A*n, t.B*n, t.C*n)
}

func (t *Triple) String() string {
	return fmt.Sprintf("(%d, %d, %d)", t.A, t.B, t.C)
}

// NewTriple ...
func NewTriple(a, b, c uint64) *Triple {
	return &Triple{a, b, c, a + b + c}
}

// PythagoreanTriplesBerggren uses Berggren's method of generating
// all primitive Pythagorean triples recursively.
// We only generate those triples where the sum of the sides is <= our limit.
func PythagoreanTriplesBerggren(limiter func(*Triple) bool) chan *Triple {
	ch := make(chan *Triple, 0)

	var generate func(t *Triple)
	generate = func(t *Triple) {
		if limiter(t) {
			return
		}
		ch <- t

		// Use the current triple as input and dive down each branch
		generate(BerggrenMatrixA(t.A, t.B, t.C))
		generate(BerggrenMatrixB(t.A, t.B, t.C))
		generate(BerggrenMatrixC(t.A, t.B, t.C))
	}

	go func() {
		generate(NewTriple(3, 4, 5))
		close(ch)
	}()
	return ch
}

// PythagoreanTriples implements incremental generation of all Pythagorean
// triples using Berggren matrices.
//
// At each level N we generate 3^N triangles. We then iterate over those results.
// Once the iterator is drained, we generate the next level of triples by
// applying the matrices to the triples from the previous level.
//
// Starts with {3,4,5} at N=1

// PythagoreanTriples ...
type PythagoreanTriples struct {
	r     []*Triple
	level int
	i     int
}

// NewPythagoreanTripleGenerator ...
func NewPythagoreanTripleGenerator() *PythagoreanTriples {
	return &PythagoreanTriples{
		r:     []*Triple{NewTriple(3, 4, 5)},
		level: 1,
		i:     -1,
	}
}

// Next ...
func (p *PythagoreanTriples) Next() bool {
	if p.i < len(p.r)-1 {
		p.i++
		return true
	}

	p.i = 0
	p.level *= 3
	r := make([]*Triple, p.level)
	j := 0
	for i := 0; i < len(p.r); i++ {
		t := p.r[i]
		r[j] = BerggrenMatrixA(t.A, t.B, t.C)
		r[j+1] = BerggrenMatrixB(t.A, t.B, t.C)
		r[j+2] = BerggrenMatrixC(t.A, t.B, t.C)
		j += 3
	}
	p.r = r
	return true
}

// Get ...
func (p *PythagoreanTriples) Get() *Triple {
	return p.r[p.i]
}

// BerggrenMatrixA is:
//  1 -2  2
//  2 -1  2
//  2 -2  3
func BerggrenMatrixA(a, b, c uint64) *Triple {
	return NewTriple(
		a-(2*b)+(2*c),
		(2*a)-b+(2*c),
		(2*a)-(2*b)+(3*c),
	)
}

// BerggrenMatrixB is:
//  −2  1  2
//  −1  2  2
//  −2  2  3
func BerggrenMatrixB(a, b, c uint64) *Triple {
	return NewTriple(
		uint64((int64(-2)*int64(a))+int64(b)+int64(2*c)),
		uint64((int64(-1)*int64(a))+int64(2*b)+int64(2*c)),
		uint64((int64(-2)*int64(a))+int64(2*b)+int64(3*c)),
	)
}

// BerggrenMatrixC is:
//  1  2  2
//  2  1  2
//  2  2  3
func BerggrenMatrixC(a, b, c uint64) *Triple {
	return NewTriple(
		a+(2*b)+(2*c),
		(2*a)+b+(2*c),
		(2*a)+(2*b)+(3*c),
	)
}
