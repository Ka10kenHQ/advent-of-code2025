package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Ka10kenHQ/advent-of-code2025/internal/day3"
	// "github.com/Ka10kenHQ/advent-of-code2025/internal/day1"
	// "github.com/Ka10kenHQ/advent-of-code2025/internal/day2"
)


func main() {
	lines := fileReader("/home/achir/dev/go/advent-of-code2025/actual.txt")

	answer1, answer2 := day3.Run(lines)
	fmt.Println(answer1, answer2)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func fileReader(abs_path string) []string {
	file, err := os.Open(abs_path)
	check(err)

	var lines []string
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
