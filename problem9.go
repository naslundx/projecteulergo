package main

import "fmt"

func Problem9() {
	for a := 1; a < 1000; a += 1 {
		for b := a; b < 1000; b += 1 {
			c := 1000 - a - b

			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}
}
