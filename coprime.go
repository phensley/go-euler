package euler

// EulersTotient computes Euler's Totient (PHI) up to N using a sieve.
func EulersTotient(max uint64) []uint64 {
	phi := make([]uint64, max+1)
	phi[1] = 1
	for n := uint64(2); n <= max; n++ {
		// The "for i" loop below handles factors of N.
		// If phi(N) == 0 right here, then N is prime.
		if phi[n] == 0 {

			// phi(prime) == prime - 1
			phi[n] = n - 1

			// Iterate over 2N, 3N, 4N ... max
			for i := n << 1; i <= max; i += n {
				// If zero, initialize to itself
				if phi[i] == 0 {
					phi[i] = i
				}
				// Incrementally calculate phi(i)
				phi[i] = phi[i] / n * (n - 1)
			}
		}
	}
	return phi
}
