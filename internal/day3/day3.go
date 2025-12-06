package day3

import (
	"math"
	"strconv"
)

func Run(lines []string) (int, int) {
	ans1 := partOne(lines)
	ans2 := partTwo(lines)

	return ans1, ans2
}

func partOne(lines []string) int {
	joltageSum := 0
	for _, line := range lines {
		joltageSum += processLine(line)
	}
	return joltageSum
}

func partTwo(lines []string) int {
	return 0
}

func processLine(line string) int {
	size := len(line)
	arr := make([]int, size)

	for i := range size {
		raw_int, err := strconv.Atoi(string(line[i]))

		if err != nil { panic("Error: Converting")
		}

		arr[i] = raw_int
	}

	curr_max := arr[0]
	second_best := math.MinInt32

	curr_max_i := 0

	for i := 1; i < size; i++ {
		if arr[i] > curr_max {
			curr_max = arr[i]
			curr_max_i = i
		}
	}

	starting_point := curr_max_i + 1

	if curr_max_i == size - 1 {
		starting_point = 0
	}

	for i := starting_point; i < size; i ++ {
		if second_best < arr[i] && second_best <= curr_max && curr_max_i != i {
			second_best = arr[i]
		}
	}

	if curr_max_i == size - 1 {
		return second_best * 10 + curr_max
	}

	return curr_max * 10 + second_best
}
