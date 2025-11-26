package main

import (
	"fmt"
)

func Problem1() {
	total := 0
	for i := 1; i < 1000; i += 1 {
		if i%3 == 0 || i%5 == 0 {
			total += i
		}
	}
	fmt.Println(total)
}
