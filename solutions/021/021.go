package euler021

import (
	"fmt"
	"math"

	"github.com/phensley/go-euler"
)

func init() {
	euler.Register("021", "Amicable numbers", solve)
}

func solve(ctx *euler.Context) {
	sums := map[int]int{}

	// Sums of divisors of all numbers < 10,000
	for n := 2; n < 10000; n++ {
		sums[n] = sumOf(divisors(n))
	}

	// Collect amicable numbers
	var amicable []int
	for n, sum := range sums {
		if n != sum && n == sums[sum] {
			amicable = append(amicable, n)
		}
	}

	// Sum of all amicable numbers
	if euler.Verbose {
		fmt.Println("Amicable numbers: ", amicable)
	}
	answer := fmt.Sprintf("%d", sumOf(amicable[0:]))
	ctx.SetAnswer(answer)
}

func divisors(n int) []int {
	sq := int(math.Sqrt(float64(n)))
	divisors := []int{1}
	for i := 2; i < sq; i++ {
		if n%i == 0 {
			divisors = append(divisors, i)
			divisors = append(divisors, n/i)
		}
	}
	// perfect square
	if n%sq == 0 {
		divisors = append(divisors, sq)
	}
	return divisors
}

func sumOf(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
