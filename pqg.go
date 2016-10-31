package pqg

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// Gen emit <p,q,g>.
// p = 2q + 1, p,q - safe primes
// g - cyclic group generator Zp
// performs n Miller-Rabin tests with 1 - 1/(4^n) probability false rate.
// Gain n - bit width for integer & probability rang for MR.
// It returns p, q, g and write error message.
func Gen(n, probability int) (*big.Int, *big.Int, *big.Int, error) {
	for {
		q, err := rand.Prime(rand.Reader, n-1)
		if err != nil {
			return nil, nil, nil, err
		}
		one := big.NewInt(1)
		two := big.NewInt(2)
		t := new(big.Int).Mul(q, two)
		p := new(big.Int).Add(t, one)
		if p.ProbablyPrime(probability) {
			for {
				g, err := rand.Int(rand.Reader, p)
				if err != nil {
					return nil, nil, nil, err
				}
				if b := new(big.Int).Exp(g, two, p); b.Cmp(one) == 0 {
					continue
				}
				if b := new(big.Int).Exp(g, q, p); b.Cmp(one) == 0 {
					return p, q, g, nil
				}
			}
		}
	}
	return nil, nil, nil, errors.New("can't emit <p,q,g>")
}
