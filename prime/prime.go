package main

import (
	"sort"
)

func generate(start int) []int {
	factors := make([]int, 0, 10)

	for candidate := 2; start > 1; candidate++ {
		for ; start%candidate == 0; start /= candidate {
			factors = append(factors, candidate)
		}

	}

	sort.Ints(factors)

	return factors
}
