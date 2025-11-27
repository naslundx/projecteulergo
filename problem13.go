package main

import (
	"bufio"
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
	scanner := bufio.NewScanner(file)
	result := BigNumber{""}
	for scanner.Scan() {
		line := scanner.Text()
		term := BigNumber{line}
		result = result.add(term)
	}

	fmt.Println(result.value[:10])
}
