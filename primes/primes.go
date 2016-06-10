// Package primes provides primitives for working with prime numbers, for
// example finding the prime factorization of a number.
package primes

// Set is a type that encapsulates a set of prime numbers and related methods.
// Typically, one will call Sieve and then Order to use this type.
type Set struct {
	set    []bool
	primes []uint64
}

// IsPrime tests if a number is prime
func (p *Set) IsPrime(x uint64) bool {
	return !p.set[x]
}

// Sieve inits a Set by finding all primes less than max.
func (p *Set) Sieve(max uint64) {
	// marking non-primes as true
	sieve := make([]bool, max)
	for i := uint64(4); i < max; i++ {
		if i%2 == 0 || i%3 == 0 {
			sieve[i] = true
		}
	}
	for p := uint64(5); p < max; p += 2 {
		if sieve[p] {
			continue
		}
		for c := p * p; c < max; c += 2 * p {
			sieve[c] = true
		}
	}
	p.set = sieve
}

// Order creates an ordered list of the primes that were sieved.
// Must be called before using factor method.
func (p *Set) Order() {
	var primes []uint64
	for idx := range p.set {
		if idx == 0 || idx == 1 {
			continue
		}
		if !p.set[idx] {
			primes = append(primes, uint64(idx))
		}
	}
	p.primes = primes
}

// Factor returns a map of the prime factors along with their exponents.
func (p *Set) Factor(x uint64) map[uint64]uint {
	factors := make(map[uint64]uint)
	for _, y := range p.primes {
		for {
			if x%y != 0 {
				break
			}
			x = x / y
			factors[y] = factors[y] + 1

		}
		if x == 1 {
			break
		}
	}
	return factors
}

// CountDivisors counts the number of divisors for x.
func (p *Set) CountDivisors(x uint64) uint {
	factors := p.Factor(x)
	count := uint(1)
	for _, v := range factors {
		count *= v + 1
	}
	return count
}

// Divisors returns a list of all the divisors for x.
func (p *Set) Divisors(x uint64) []uint64 {
	factors := p.Factor(x)
	divs := []uint64{1}
	for p, e := range factors {
		y := uint64(1)
		for i := uint(0); i < e; i++ {
			y *= p
			for _, z := range divs {
				divs = append(divs, p*z)
			}
			divs = append(divs, p)
		}
	}
	return divs
}
