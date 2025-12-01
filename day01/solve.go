package day01

import (
	"bufio"
	"strconv"
	"strings"
)

func SolvePuzzle1(input string) string {
	position := 50
	resetOccurences := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		// First char is direction
		dir := string(line[0])
		// Rest is distance
		distance, _ := strconv.Atoi(line[1:])
		position = TurnDial(dir, distance, position)

		if position == 0 {
			resetOccurences++
		}
	}
	return strconv.Itoa(resetOccurences)
}

func TurnDial(direction string, distance int, startPosition int) int {
	endPosition := calculateEndPosition(direction, distance, startPosition)
	return endPosition
}

func TurnDialWithCrossings(direction string, distance int, startPosition int) (int, int) {
	var zeroClicks int

	if direction == "R" {
		firstZero := (100 - startPosition) % 100
		if firstZero == 0 {
			firstZero = 100
		}
		if firstZero <= distance {
			zeroClicks = 1 + (distance-firstZero)/100
		}
	} else if direction == "L" {
		// First 0 is hit after startPosition clicks (if startPosition > 0)
		// or after 100 clicks (if startPosition == 0)
		firstZero := startPosition
		if firstZero == 0 {
			firstZero = 100
		}
		if firstZero <= distance {
			zeroClicks = 1 + (distance-firstZero)/100
		}
	}

	endPosition := calculateEndPosition(direction, distance, startPosition)
	return endPosition, zeroClicks
}

func calculateEndPosition(direction string, distance int, startPosition int) int {
	var endPosition int
	if direction == "R" {
		endPosition = startPosition + distance
	} else if direction == "L" {
		endPosition = startPosition - distance
	}

	endPosition = endPosition % 100
	if endPosition < 0 {
		endPosition += 100
	}
	return endPosition
}

func SolvePuzzle2(input string) string {
	position := 50
	zeroCrossings := 0
	scanner := bufio.NewScanner(strings.NewReader(input))
	for scanner.Scan() {
		line := scanner.Text()
		dir := string(line[0])
		distance, _ := strconv.Atoi(line[1:])

		var crossings int
		position, crossings = TurnDialWithCrossings(dir, distance, position)
		zeroCrossings += crossings
		// fmt.Printf("%s: pos %d -> %d, clicks on 0: %d, total: %d\n", line, position, position, crossings, zeroCrossings)
	}
	return strconv.Itoa(zeroCrossings)
}
