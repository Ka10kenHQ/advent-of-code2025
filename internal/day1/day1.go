package day1

import (
	"strconv"
)

const STARTING_POSITION = 50

func Run(lines []string) (int, int) {
	answer1 := partOne(lines)
	answer2 := partTwo(lines)

	return answer1, answer2
}

func partOne(lines []string) int {
	current_position := STARTING_POSITION
	zero_count := 0
	for _, line := range lines {
		direction, amount := processLine(line)
		current_position, _ = calculateNextPosition(direction, amount, current_position)

		if current_position == 0 {
			zero_count++
		}
	}
	return zero_count
}

func partTwo(lines []string) int {
	current_position := STARTING_POSITION
	zero_count := 0
	cnt := 0
	for _, line := range lines {
		direction, amount := processLine(line)
		current_position, cnt = calculateNextPosition(direction, amount, current_position)
		zero_count += cnt
	}
	return zero_count
}

func calculateNextPosition(direction string, amount, currentPosition int) (int, int) {
	switch direction {
	case "L":
		cnt := rotateLeft(currentPosition, amount)
		currentPosition = mod100(currentPosition - amount)
		return currentPosition, cnt
	case "R":
		cnt := rotateRight(currentPosition, amount)
		currentPosition = mod100(currentPosition + amount)
		return currentPosition, cnt
	default:
		panic("Invalid direction")
	}
}

func rotateLeft(currentPosition, amount int) int {
	count := amount / 100
	steps := amount % 100
	at := currentPosition

	if at == 0 {
		at = 100
	}

	if at <= steps {
		count++
	}
	return count
}

func rotateRight(currentPosition, amount int) int {
	count := amount / 100
	steps := amount % 100

	if currentPosition+steps >= 100 {
		count++
	}
	return count
}

func mod100(x int) int {
	return ((x % 100) + 100) % 100
}

func processLine(line string) (string, int) {
	direction := string(line[0])
	amount, err := strconv.ParseInt(line[1:], 10, 32)
	if err != nil {
		panic(err)
	}
	return direction, int(amount)
}
