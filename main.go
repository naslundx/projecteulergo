package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <problem_id>")
		return
	}

	id, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("Error: problem_id must be an integer")
		return
	}

	problems := map[int]func(){
		0:  Problem0,
		1:  Problem1,
		2:  Problem2,
		3:  Problem3,
		4:  Problem4,
		6:  Problem6,
		7:  Problem7,
		8:  Problem8,
		9:  Problem9,
		13: Problem13,
		14: Problem14,
		25: Problem25,
	}

	run, ok := problems[id]
	if !ok {
		fmt.Printf("Error: problem %d not found.\n", id)
		return
	}

	run()
}
