package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(path string) string {
	b, _ := os.ReadFile(path)
	return strings.TrimSpace(string(b))
}

func ReadLines(path string) []string {
	input := ReadInput(path)
	return strings.Split(strings.TrimSpace(input), "\n")
}

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
