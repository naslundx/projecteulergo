package main

import "fmt"

func Problem25() {
	a := BigNumber{"1"}
	b := BigNumber{"1"}
	index := 2

	for b.digits() < 1000 {
		index += 1
		a, b = b, a.add(b)
	}

	fmt.Println(index)

}
