package main

import (
	"fmt"
	"io"
	"os"
)

func Problem8() {
	N := 13

	file, err := os.Open("data/problem8.txt")
	if err != nil {
		fmt.Println("Could not open file.")
		return
	}
	defer file.Close()

	data, _ := io.ReadAll(file)
	str := string(data)
	best := 0

	for i, _ := range str {
		total := 1
		for j := i; j < i+N && j < len(str); j += 1 {
			total *= (int(str[j]) - 48)
		}
		if total > best {
			best = total
		}
	}

	fmt.Println(best)
}
