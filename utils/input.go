package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(path string) (string, error) {
	b, err := os.ReadFile(path)
	return string(b), err
}

func ReadLines(path string) []string {
	input, err := ReadInput(path)
	CheckErr(err)
	return strings.Split(strings.TrimSpace(input), "\n")
}

func CheckErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
