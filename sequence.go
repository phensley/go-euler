package euler

import "math/big"

// Fibonacci holds the current state of a Fibonacci sequence
type Fibonacci struct {
	a uint64
	b uint64
	i uint
}

// NewFibonacci constructs a new Fibonacci sequence generator
func NewFibonacci(a, b uint64) *Fibonacci {
	return &Fibonacci{a, b, 0}
}

// Next returns the next value in the Fibonacci sequence
func (f *Fibonacci) Next() uint64 {
	r := uint64(0)
	switch f.i {
	case 0:
		r = f.a
	case 1:
		r = f.b
	default:
		f.a, f.b = f.b, f.a+f.b
		r = f.b
	}
	f.i++
	return r
}

// FibonacciSequence generates the fibonacci sequence into a channel
func FibonacciSequence(a, b uint64) <-chan uint64 {
	ch := make(chan uint64, 0)
	f := NewFibonacci(a, b)
	go func() {
		for {
			ch <- f.Next()
		}
	}()
	return ch
}

// FibonacciBig holds the current state of a Fibonacci sequence as
// calculated using big.Int instances
type FibonacciBig struct {
	a *big.Int
	b *big.Int
	i uint
}

// NewFibonacciBig constructs a new FibonacciBig sequence generator
func NewFibonacciBig(a, b int64) *FibonacciBig {
	return &FibonacciBig{big.NewInt(a), big.NewInt(b), 0}
}

// Next updates the argument with the next value in the Fibonacci sequence
func (f *FibonacciBig) Next(n *big.Int) {
	switch f.i {
	case 0:
		n.Set(f.a)
	case 1:
		n.Set(f.b)
	default:
		n.Add(f.a, f.b)
		f.a.Set(f.b)
		f.b.Set(n)
	}
	f.i++
}

// FibonacciBigSequence returns a channel which generates the FibonacciBig sequence
func FibonacciBigSequence(a, b int64) <-chan *big.Int {
	ch := make(chan *big.Int, 0)
	go func() {
		f := NewFibonacciBig(a, b)
		n0 := &big.Int{}
		n1 := &big.Int{}
		for {
			f.Next(n0)
			// Swap instances to avoid immediately updating the reference we're
			// about to pass through the channel
			n1, n0 = n0, n1
			ch <- n1
		}
	}()
	return ch
}
