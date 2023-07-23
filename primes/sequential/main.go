package main

import (
	"fmt"
)

type PrimeResult struct {
	n     int
	prime bool
}

var Sample = []PrimeResult{
	{2, true},
	{142702110479723, true},
	{299593572317531, true},
	{3333333333333301, true},
	{3333333333333333, false},
	{4444444444444423, true},
	{4444444444444444, false},
	{4444444488888889, false},
	{5555553133149889, false},
	{5555555555555503, true},
	{5555555555555555, false},
	{6666666666666666, false},
	{6666666666666719, true},
	{6666667141414921, false},
	{7777777536340681, false},
	{7777777777777753, true},
	{7777777777777777, false},
	{9999999999999917, true},
	{9999999999999999, false},
}

func main() {
	for _, fixture := range Sample {
		primeRes := IsPrime(fixture.n)
		if primeRes != fixture.prime {
			panic(fmt.Sprintf("Assertion failed for %v", fixture.n))
		}
		label := ""
		if fixture.prime {
			label = "prime"
		} else {
			label = ""
		}
		fmt.Printf("%v\t%v\n", label, fixture.n)
	}
}
