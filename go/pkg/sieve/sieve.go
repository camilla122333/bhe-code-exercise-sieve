package sieve

import (
	"math"
)

// maxPrimeIndex is the largest prime index from the required tests
const maxPrimeIndex = 10000000

type eratosthenesSieve struct {
	primes []int64
}

// Sieve provides a method for retrieving prime numbers by index
type Sieve interface {
	NthPrime(n int64) int64
}

// NewSieve creates and returns a Sieve implementation, precomputing primes once
// up to maxPrimeIndex. This ensures performace for repeated calls to NthPrime
func NewSieve() Sieve {
	p := findPrimes(maxPrimeIndex)
	return &eratosthenesSieve{primes: p}
}

// NthPrime returns the nth prime number
func (s *eratosthenesSieve) NthPrime(n int64) int64 {
	if n < 0 {
		panic("Invalid input: negative index not allowed")
	}

	return s.primes[n]
}

// returns a list of primes based on a max prime index
func findPrimes(n int64) []int64 {
	p := []int64{}

	// calculate the limit using the Prime Number Theorem:
	// to later return a list of primes, we need to figure out a safe upper limit
	// in order to have a safe range of numbers to sift through.
	// formula: limit = n * (logn + loglogn)

	// add 1 to adjust the max number to 0-based indexing (formula is 1-based)
	n += 1
	logN := math.Log(float64(n))
	logLogN := math.Log(logN)
	limitFl := math.Ceil(float64(n) * (logN + logLogN))
	limit := int64(limitFl)

	// eratosthenes algorithm
	// initialize numbers as potentially prime (true)
	marked := make([]bool, limit+1)
	for i := 2; i < len(marked); i++ {
		marked[i] = true
	}

	// mark all composite (non-prime) numbers as false
	for i := 2; i*i < len(marked); i++ {
		if marked[i] {
			smallPrime := i
			for j := smallPrime * smallPrime; j < len(marked); j += smallPrime {
				marked[j] = false
			}
		}
	}

	// collect all numbers that remain marked as prime
	for i := range marked {
		if marked[i] {
			p = append(p, int64(i))
		}
	}

	return p
}
