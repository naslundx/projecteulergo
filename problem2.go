package main

import (
	"fmt"
)

func Problem2() {
	a, b := 0, 1
	result := 0

	for b <= 4000000 {
		if b%2 == 0 {
			result += b
		}
		a, b = b, a+b
	}

	fmt.Println(result)
}
