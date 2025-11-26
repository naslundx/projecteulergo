package main

import (
	"fmt"
)

var results = make(map[int]int)

func Collatz(value int) int {
	steps := 0
	for {
		if value == 1 {
			return steps
		}

		if val, ok := results[value]; ok {
			return steps + val
		}

		steps += 1
		if value%2 == 0 {
			value = value / 2
		} else {
			value = value*3 + 1
		}
	}
	return -1
}

func Problem14() {
	bestIndex, bestValue := 0, 0
	results[1] = 1
	results[2] = 2

	for i := 3; i < 1000000; i += 1 {
		result := Collatz(i)
		results[i] = result
		if result > bestValue {
			bestIndex, bestValue = i, result
		}
	}

	fmt.Println(bestIndex)
}
