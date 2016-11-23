package euler072

import (
	"fmt"
	"math/big"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("072", "Counting fractions", solve)
}

func solve(ctx *euler.Context) {

	// Brute-force of the known case clarifies the relationship to Euler's Totient
	if euler.Verbose {
		known()
	}

	// The number of reduced fractions up to d <= 1,000,000 is related to the
	// count of coprime numbers. Reducing fractions by dividing both the numerator
	// and denominator by the greatest common factor between them.
	//
	// Coprimality is based on gcf(a, b) as well, so the count of distinct
	// reduced fractions up to N is related to the count of coprime numbers <= N,
	// which is also known as Euler's Totient or phi(N).  To get our answer we
	// just need to sum the phi(a) series for 2 <= a <= N.

	d := uint64(1000000)
	sum := uint64(0)
	for n, phi := range euler.EulersTotient(d) {
		// Skip 0 and 1 since we don't consider fractions n/0 or n/1
		if n >= 2 {
			sum += phi
		}
	}

	answer := fmt.Sprintf("%d", sum)
	ctx.SetAnswer(answer)
}

func known() {
	fmt.Print("\nComputing known case for N=8\n\n")
	num := int64(8)
	seq := []int64{}
	p := int64(0)
	for n := int64(2); n <= num; n++ {
		c := calculate(n)
		d := c - p
		seq = append(seq, d)
		p = c
		fmt.Printf("N=%d count=%d\n", n, c)
	}

	fmt.Println("Deltas between successive counts:", seq)

	fmt.Printf("\nEuler phi(n) for 2 <= n <= %d: ", num)
	sum := uint64(0)
	for n, phi := range euler.EulersTotient(uint64(num)) {
		if n >= 2 {
			sum += phi
			fmt.Print(phi, " ")
		}
	}
	fmt.Println("\nSum: ", sum)
}

func calculate(lim int64) int64 {
	count := int64(0)
	seen := make(map[string]*struct{})

	for d := int64(2); d <= lim; d++ {
		for n := int64(1); n < d; n++ {
			r := big.NewRat(n, d)
			s := r.String()
			if seen[s] == nil {
				seen[s] = &struct{}{}
				count++
			}
		}
	}
	return count
}
