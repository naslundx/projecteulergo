package main

import (
	"fmt"
	"math"
)

func Reverse(N int) (result int) {
	length := int(math.Log10(float64(N)))

	for length >= 0 {
		length -= 1
		n := N % 10
		result = result*10 + n
		N /= 10
	}
	return
}

func Problem4() {
	largest := 0
	for i := 100; i < 1000; i += 1 {
		for j := i; j < 1000; j += 1 {
			N := i * j
			if N == Reverse(N) && N > largest {
				largest = N
			}
		}
	}
	fmt.Println(largest)
}
