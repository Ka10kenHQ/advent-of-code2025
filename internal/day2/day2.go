package day2

import (
	"strconv"
	"strings"

	"github.com/go-playground/locales/bo"
)

func Run(lines []string) (int64, int) {
	ans1 := partOne(lines)
	return ans1, 0
}

func partOne(lines [] string) int64 {
	var sum int64 = 0

	for _, line := range lines {
		sum += processLine(line)
	}
	return sum
}

func processLine(line string) int64 {
	var invalidNumbersSum int64 = 0
	rangeList := strings.SplitSeq(line, ",")

	for item := range rangeList {
		ran := strings.Split(item, "-")
		start, err1 := strconv.ParseInt(ran[0], 10, 64)
		end, err2 := strconv.ParseInt(ran[1], 10, 64)
		
		if err1 != nil || err2 != nil {
			panic("Failed to convert ranges to ints")
		}

		checkedRange := checkValidity(start, end, isValidIdPart1)
		for _, invalidNumber := range checkedRange {
			invalidNumbersSum += invalidNumber
		}
	}

	return invalidNumbersSum
}

func isValidIdPart1(numStr string) bool {
	strLength := len(numStr)

	if mod(strLength, 2) != 0 {
		return true
	}

	return !(numStr[:strLength/2] == numStr[strLength/2:])
}

func isValidIdPart2(numStr string) bool {
	// part two coming soon
	return true
}

func checkValidity(start, end int64, predicate func(string) bool) []int64 {
	var result []int64

	for i := start; i <= end; i ++ {
		numStr := strconv.FormatInt(i, 10)

		if !predicate(numStr) {
			result = append(result, i)
		}
	}

	return result
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}
