package main

import (
	"fmt"
)

var results = make(map[int]int)

func Collatz(value int) int {
	if value == 1 {
		return 0
	}

	if val, ok := results[value]; ok {
		return val
	}

	if value%2 == 0 {
		value /= 2
	} else {
		value = value*3 + 1
	}

	return 1 + Collatz(value)
}

func Problem14() {
	bestIndex, bestValue := 0, 0

	for i := 1; i < 1000000; i += 1 {
		result := Collatz(i)
		results[i] = result
		if result > bestValue {
			bestIndex, bestValue = i, result
		}
	}

	fmt.Println(bestIndex)
}
