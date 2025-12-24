package main

import (
	"fmt"
	"strconv"

	"github.com/sssdrum/aoc-2025/utils"
)

const (
	START = 0
	END   = 99
)

func main() {
	part1()
}

func part1() {
	curr := 50
	res := 0
	lines := utils.ReadLines("./input.txt")
	for _, l := range lines {
		dir, tok := l[0], l[1:]
		steps, err := strconv.Atoi(tok)
		utils.CheckErr(err)

		if dir == 'R' {
			curr = (curr + steps) % (END + 1)
		} else {
			curr = ((curr - steps) + END + 1) % (END + 1)
		}

		if curr == 0 {
			res++
		}
	}

	fmt.Println(res)
}
