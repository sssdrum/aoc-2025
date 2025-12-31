package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/sssdrum/aoc-2025/utils"
)

func main() {
	part1()
}

func part1() {
	lines := utils.ReadLines("./input.txt")
	ops := strings.Fields(lines[len(lines)-1])
	tots := make([]int, len(ops))
	for i, o := range ops {
		if o == "*" {
			tots[i] = 1
		}
	}

	for _, l := range lines[:len(lines)-1] {
		fields := strings.Fields(l)
		for i, f := range fields {
			n, _ := strconv.Atoi(f)
			if ops[i] == "*" {
				tots[i] *= n
			} else {
				tots[i] += n
			}
		}
	}

	res := 0
	for _, t := range tots {
		res += t
	}
	fmt.Println(res)
}
