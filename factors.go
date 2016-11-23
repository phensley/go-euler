package euler

import "math"

// ProperDivisorSieve ...
func ProperDivisorSieve(n int) [][]int {
	r := make([][]int, n)
	for i := 0; i < n; i++ {
		r[i] = []int{1}
	}
	for i := 2; i < n/2; i++ {
		for j := 2 * i; j < n; j += i {
			r[j] = append(r[j], i)
		}
	}
	return r
}

// ProperDivisorSumSieve ..
func ProperDivisorSumSieve(n int) []int {
	r := make([]int, n)
	for i := 0; i < n; i++ {
		r[i] = 1
	}
	for i := 2; i < n/2; i++ {
		for j := 2 * i; j < n; j += i {
			r[j] += i
		}
	}
	return r
}

// GreatestCommonDivisor of N and D
func GreatestCommonDivisor(n, d uint64) uint64 {
	for d != 0 {
		t := n % d
		n, d = d, t
	}
	return n
}

// EuclidGCD returns the greatest common divisor of N and D
func EuclidGCD(n, d uint64) uint64 {
	for n != d {
		if n > d {
			n = n - d
		}
		if d > n {
			d = d - n
		}
	}
	return n
}

// ProperDivisors ...
func ProperDivisors(n int) []int {
	res := []int{1}
	lim := int(math.Floor(math.Sqrt(float64(n))))
	for i := 2; i <= lim; i++ {
		if n%i == 0 {
			res = append(res, i)
			v := n / i
			if v > lim {
				res = append(res, v)
			}
		}
	}
	return res
}

// SumOfProperDivisors returns the sum of the proper divisors
func SumOfProperDivisors(n int) int {
	lim := int(math.Floor(math.Sqrt(float64(n))))
	sum := 1
	for i := 2; i <= lim; i++ {
		if n%i == 0 {
			sum += i
			v := n / i
			if v > lim {
				sum += v
			}
		}
	}
	return sum
}
