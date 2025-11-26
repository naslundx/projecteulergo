package main

import (
	"fmt"
)

func Problem6() {
	N := 100

	sumOfSquares := 0
	squareOfSum := 0

	for i := 1; i <= N; i += 1 {
		sumOfSquares += i * i
		squareOfSum += i
	}

	squareOfSum *= squareOfSum

	fmt.Println(squareOfSum - sumOfSquares)
}
