package main

import (
	"fmt"
	"os"
)

func Problem13() {
	file, err := os.Open("data/problem13.txt")
	if err != nil {
		fmt.Println("Could not open file.")
		return
	}
	defer file.Close()

	// in progress
}
