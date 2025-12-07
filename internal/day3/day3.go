package day3

import (
	"fmt"
	"strconv"
)

func Run(lines []string) (int64, int64) {
	ans1 := partOne(lines)
	ans2 := partTwo(lines)

	return ans1, ans2
}

func partOne(lines []string) int64 {
	var joltageSum int64 = 0
	for _, line := range lines {
		joltageSum += processLine(line, 2)
	}
	return joltageSum
}

func partTwo(lines []string) int64 {
	var joltageSum int64 = 0
	for _, line := range lines {
		joltageSum += processLine(line, 12)
	}
	return joltageSum
}

func processLine(line string, n int64) int64 {
	var size int64 = int64(len(line))
	arr := make([]int64, size)

	for i := range size {
		raw_int, err := strconv.Atoi(string(line[i]))
		if err != nil { panic("Error: Converting") }
		arr[i] = int64(raw_int)
	}
	
	result := make([]int64, 0, n)
	var numToSkip int64 = size - n
	
	for i := range size {
		for len(result) > 0 && result[len(result)-1] < arr[i] && numToSkip > 0 {
			result = result[:len(result)-1]
			numToSkip--
		}
		result = append(result, arr[i])
	}
	
	result = result[:n]
	
	fmt.Println(result)
	ans := buildAnswer(result, n)
	fmt.Println(ans)
	return ans
}


func buildAnswer(digits []int64, n int64) int64 {
	var ans int64 = 0
	for i := range n {
		power := n - 1 - i
		ans += digits[i] * pow10(power)
	}
	return ans
}

func pow10(n int64) int64 {
	var result int64 = 1
	for range n {
		result *= 10
	}
	return result
}
