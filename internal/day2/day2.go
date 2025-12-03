package day2

import (
	"strconv"
	"strings"
)

func Run(lines []string) (int64, int64) {
	ans1 := partOne(lines, isValidIdPart1)
	ans2 := partOne(lines, isValidIdPart2)
	return ans1, ans2
}

func partOne(lines [] string, predicate func(string) bool) int64 {
	var sum int64 = 0

	for _, line := range lines {
		sum += processLine(line, predicate)
	}
	return sum
}

func processLine(line string, predicate func(string) bool) int64 {
	var invalidNumbersSum int64 = 0
	rangeList := strings.SplitSeq(line, ",")

	for item := range rangeList {
		ran := strings.Split(item, "-")
		start, err1 := strconv.ParseInt(ran[0], 10, 64)
		end, err2 := strconv.ParseInt(ran[1], 10, 64)
		
		if err1 != nil || err2 != nil {
			panic("Failed to convert ranges to ints")
		}

		checkedRange := checkValidity(start, end, predicate)
		for _, invalidNumber := range checkedRange {
			invalidNumbersSum += invalidNumber
		}
	}

	return invalidNumbersSum
}

func isValidIdPart1(numStr string) bool {
	strLength := len(numStr)

	if mod(strLength, 2) != 0 {
		return false
	}

	return numStr[:strLength/2] == numStr[strLength/2:]
}

func isValidIdPart2(numStr string) bool {
	strLen := len(numStr)

	for i := 1; i < strLen; i++ {

		if strLen % i != 0 {
			continue
		}

		subStr := numStr[:i]
		j := 0

		for j < strLen && j + i <= strLen {
			if subStr != numStr[j:j+i] {
				break
			}
			j+=i
		}

		if j == strLen {
			return true
		}
	}

	return false
}

func checkValidity(start, end int64, predicate func(string) bool) []int64 {
	var result []int64

	for i := start; i <= end; i ++ {
		numStr := strconv.FormatInt(i, 10)

		if predicate(numStr) {
			result = append(result, i)
		}
	}

	return result
}

func mod(a, b int) int {
	return ((a % b) + b) % b
}
