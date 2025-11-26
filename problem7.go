package main

import "fmt"

func Problem7() {
	const N = 10001
	primes := make([]int, 0, N+1)
	primes = append(primes, 2)

	for n := 3; len(primes) < N; n += 2 {
		divisible := false
		for _, p := range primes {
			if p*p > n {
				break
			}

			if n%p == 0 {
				divisible = true
				break
			}
		}

		if !divisible {
			primes = append(primes, n)
		}
	}

	fmt.Println(primes[N-1])
}
