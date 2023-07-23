package main

import "math"

func IsPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	root := int(math.Sqrt(float64(n)))
	for i := 3; i <= root; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
