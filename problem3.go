package main

import (
	"fmt"
)

func Problem3() {
	n := 600851475143

	for {
		fmt.Println("iteration")
		found := false
		for i := 2; i < n/2; i += 1 {
			if n%i == 0 {
				n /= i
				fmt.Println(i)
				found = true
				break
			}
		}
		if !found {
			fmt.Println(n)
			return
		}
	}
}
