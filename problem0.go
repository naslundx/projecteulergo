package main

import (
	"fmt"
)

func Problem0() {
	total := 0
	for i := 1; i <= 173000; i += 1 {
		if i%2 == 1 {
			total += i * i
		}
	}
	fmt.Println(total)
}
