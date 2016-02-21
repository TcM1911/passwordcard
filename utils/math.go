// Package utils is a collection of utility functions for PasswordCard.
package utils

// MathMod calculates the n mod m.
func MathMod(n, m int) int {
	var result int
	if n%m == 0 {
		result = 0
	} else if n >= 0 {
		result = n % m
	} else {
		result = m + (n % m)
	}
	return result
}

// Randomer is an pseudorandom number generator.
type Randomer interface {
	nextInt(i int) int
}
